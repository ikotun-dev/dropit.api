package controllers

import (
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
		w.Write([]byte("Session created successfully"))

	} else {
		w.WriteHeader(http.StatusBadRequest)
		res := "session key must be greater than 6 characters"
		w.Write([]byte(res))

	}
}
