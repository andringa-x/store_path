package api

import (
	"encoding/json"
	"net/http"
)

type StorePathBody struct {
	Path []string `json:"path"`
	List string   `json:"list"`
}

type StorePathReponse struct {
	Code int
	List string
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred.", http.StatusInternalServerError)
	}

	GeminiErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An error occurred while contacting Gemini.", http.StatusInternalServerError)
	}
)
