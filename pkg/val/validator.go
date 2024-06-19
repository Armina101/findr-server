package val

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translate "github.com/go-playground/validator/v10/translations/en"
)

var timeOut = 10 * time.Second

var errs []error

// FieldValidator create a contextual validation of all the struct field and their struct
// tags and proper configuration of the validation error
func FieldValidator(s interface{}) []error {
	vald := validator.New()
	ctx, cancelCtx := context.WithTimeout(context.Background(), timeOut)
	defer cancelCtx()

	// setting english translator
	eng := en.New()
	translate := ut.New(eng, eng)
	translator, ok := translate.GetTranslator("en")

	if !ok {
		fmt.Println("unable to set english translator")
		return []error{fmt.Errorf("cannot process validation of input")}
	}

	err := en_translate.RegisterDefaultTranslations(vald, translator)
	if err != nil {
		return []error{fmt.Errorf("cannot register default translation")}
	}

	err = vald.StructCtx(ctx, s)
	var fieldErrors validator.ValidationErrors
	errors.As(err, &fieldErrors)

	for _, v := range fieldErrors {
		err := fmt.Errorf(v.Translate(translator))
		errs = append(errs, err)
	}

	return errs
}

