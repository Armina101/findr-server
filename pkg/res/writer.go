package res

import (
	"encoding/json"
	"net/http"
)

// Writer this function writes back a response to the client after processing a request
func Writer(wr http.ResponseWriter, statusCode int, resp map[string]interface{}) error {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(statusCode)

	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = wr.Write(b)
	if err != nil {
		return err
	}
	return nil
}
