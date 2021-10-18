package http_utils

import (
	"encoding/json"
	"github.com/Moriartii/bookstore_items-api/utils/errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err errors.RestErr) {
	RespondJson(w, err.Status, err)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(err.Status)
	// json.NewEncoder(w).Encode(err)
}
