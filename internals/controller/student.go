package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/thebravebyte/findr/internals"
	"github.com/thebravebyte/findr/internals/auth"
	"github.com/thebravebyte/findr/pkg/enc"
	"github.com/thebravebyte/findr/pkg/req"
	"github.com/thebravebyte/findr/pkg/res"
)

// Todo: Endpoints implementation start here

func (fr *FindrApp) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var std internals.Students
		
		err := req.ReadAndValidate(w, r, &std); if err != nil{
			res.Writer(w, http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		// encrypt the students password using argon21d package
		password, err := enc.CreateHash(std.Password)
		if err != nil {
			fr.Logger.Error(err.Error())
			_ = res.Writer(w, http.StatusBadRequest, map[string]interface{}{
				"error":   err.Error(),
				"message": "unable to secure your credentials",
			})
			return
		}

		std.Password = password
		std.UpdatedAt, std.CreatedAt = time.Now(), time.Now()

		// database queries to add new user
		ok, check, err := fr.DB.CreateAccount(std)

		if err != nil {
			_ = res.Writer(w, http.StatusInternalServerError, map[string]interface{}{
				"error":   err,
				"message": fmt.Sprintf("unable to add this user - %v", std.Email),
			})
			return
		}

		// checks if the user already exists or not and provides appropriate JSON responses
		switch ok && check {
		case true:
			// returns a JSON response indicating that the account already exists
			_ = res.Writer(w, http.StatusConflict, map[string]interface{}{
				"error":   err,
				"message": fmt.Sprintf("This account already exist,%v", std.Email),
			})

			return

		case false:
			// sends an email notification and returns a JSON response indicating successful account creation
			subName := fmt.Sprintf("%s %s", std.FirstName, std.LastName)
			mail := internals.Mail{
				Source:      os.Getenv("SOURCE_EMAIL_ADDRESS"),
				Destination: std.Email,
				Name:        subName,
				Message: fmt.Sprintf(`
				<h2>Welcome to Findr </h2>
				<p>Hello %s,</p>
				<p>Your account has been successfully created.</p>`, subName),
				Subject:  "Findr Account Creation",
				Template: "template.html",
			}

			// sending the mail through the mail channels
			fr.MailChan <- mail

			_ = res.Writer(w, http.StatusInternalServerError, map[string]interface{}{
				"message": "Account created Successfully",
			})
			return

		default:
			// handles unexpected scenarios by returning a JSON response with an appropriate error message
			_ = res.Writer(w, http.StatusInternalServerError, map[string]interface{}{
				"error":   err,
				"message": "Something went wrong",
			})
			return
		}

	}
}

// SignInJWT this is to process the authorized access for user by generating validating the password via
// a database and using jwt token code to authorize the user access

type LoginStudent struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (fr *FindrApp) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var stdLg LoginStudent
		var std internals.Students

		//	get the request body, parse and validate all the data
		err := req.ReadAndValidate(w, r, &stdLg); if err != nil{
			res.Writer(w, http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}


		//	query the database to verify a std account and details
		verifyStd, err := fr.DB.VerifyDetails(stdLg.Email)
		if err != nil {
			res.Writer(w, http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		std = verifyStd
		// verify the password with the hashed password
		ok, err := enc.VerifyPassword(std.Password, stdLg.Password)
		if err != nil && !ok {
			res.Writer(w, http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return

		} else {

			token, err := auth.CreateToken(std.ID, stdLg.Email, stdLg.Password, fr.Logger)
			if err != nil {
				res.Writer(w, http.StatusInternalServerError, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			// store the token to the database
			ok, err := fr.DB.UpdateDetails(std.ID, token)
			if err != nil && !ok {
				res.Writer(w, http.StatusInternalServerError, map[string]interface{}{
					"error": fmt.Sprintf("Update: error this query : %s", err.Error()),
				})
				return
			}

			res.Writer(w, http.StatusOK, map[string]interface{}{
				"email":      stdLg.Email,
				"std_id":     std.ID,
				"first_name": std.FirstName,
				"last_name":  std.LastName,
				"auth_token": token,
			})
		}
	}
}


func (fr *FindrApp) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
