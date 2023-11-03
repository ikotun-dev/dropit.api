package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ikotun-dev/clipsync/pkg/helpers"
	"github.com/ikotun-dev/clipsync/pkg/middleware"
	"github.com/ikotun-dev/clipsync/pkg/models"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	//session_key = vars["session_key"]

	SessionToCreate := &models.Session{}
	helpers.ParseBody(r, SessionToCreate)
	//fmt.Print(r)
	//if the session key is less than 7 characters
	if len(SessionToCreate.Session_key) < 6 {
		err := SessionToCreate.CreateSession()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the session ID.
		fmt.Println(SessionToCreate.Session_key)

		// Generate a JWT token with the session key
		token, err := middleware.CreateJWT(SessionToCreate.Session_key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		res := map[string]string{"message": "Session created successfully", "token": token}
		response, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Authorization", "Bearer "+token)
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
