package controllers

import (
	"net/http"

	"github.com/ikotun-dev/clipsync/pkg/helpers"
	"github.com/ikotun-dev/clipsync/pkg/models"
)

var NewSession models.Session

func CreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//vars := mux.Vars(r)
	//session_key = vars["session_key"]
	SessionToCreate := &models.Session{}
	helpers.ParseBody(r, SessionToCreate)
	if len(SessionToCreate.Session_key) > 8 {

	} else {
		w.WriteHeader(http.StatusBadRequest)
		res := "session key must be greater than 6 characters"
		w.Write([]byte(res))

	}

}
