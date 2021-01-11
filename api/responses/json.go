package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON performs the formatting of a JSON response using the passed in
// HTTP response writer.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR uses the JSON function to format error responses based
// on the err parameter and writes to the HTTP response writer.
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
