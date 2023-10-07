package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ikotun-dev/clipsync/pkg/helpers"
	"github.com/ikotun-dev/clipsync/pkg/models"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//vars := mux.Vars(r)
	//session_key = vars["session_key"]
	SessionToCreate := &models.Session{}
	helpers.ParseBody(r, SessionToCreate)

	//if the session key is less than 7 characters
	if len(SessionToCreate.Session_key) > 7 {
		err := SessionToCreate.CreateSession()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		res := map[string]string{"message": "Session created successfully"}
		response, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(response)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]string{"error": "session key should not be less than 7 characters"}
		errorMessage, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(errorMessage))

	}
}
