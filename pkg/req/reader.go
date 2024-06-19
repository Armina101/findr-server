package req

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/thebravebyte/findr/pkg/val"
)

// Reader this function helps to read from the request body
func Reader(w http.ResponseWriter, r *http.Request, data interface{}) (interface{}, error) {
	rd := http.MaxBytesReader(w, r.Body, int64(1024*1024)*2)
	defer func(rd io.ReadCloser) {
		err := rd.Close()
		if err != nil {
			panic(err)
		}
	}(rd)

	// decode the encoded data from the request body
	if err := json.NewDecoder(rd).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// ReadAndValidate this function helps to read and validate the request body using the functions define above
func ReadAndValidate(w http.ResponseWriter, r *http.Request, d interface{}) error {
	res, err := Reader(w, r, &d)
	if err != nil {
		fmt.Printf("ReadAndValidateError: %s", err.Error())
		http.Error(w, "unable to read the balance request body", http.StatusBadRequest)
	}
	if errs := val.FieldValidator(&res); errs != nil {
		for _, e := range errs {
			fmt.Printf("ReadAndValidateError: %s", e.Error())
		}

		http.Error(w, "error while validating user invalid inputs", http.StatusBadRequest)
		return errors.New("invalid input from user")
	}
	return nil
}
