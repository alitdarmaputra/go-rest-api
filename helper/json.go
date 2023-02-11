package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	err := json.NewDecoder(r.Body).Decode(&result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, responses interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responses)
	PanicIfError(err)
}
