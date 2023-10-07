package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// function to join session
func JoinSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	session_key := vars["session_key"]
	if len(session_key) < 7 {
		res := map[string]string{"error": "session key should not be less than 7 characters"}
		errorMessage, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(errorMessage))

	}

}
