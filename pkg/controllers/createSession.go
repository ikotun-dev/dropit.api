package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ikotun-dev/clipsync/pkg/helpers"
	"github.com/ikotun-dev/clipsync/pkg/middleware"
	"github.com/ikotun-dev/clipsync/pkg/models"
)

var sessionWebSocketMap = make(map[string]*websocket.Conn)

func CreateSession(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
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
		// Generate a JWT token with the session key
		token, err := middleware.CreateJWT(SessionToCreate.Session_key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Store the WebSocket connection associated with the session_key
		sessionWebSocketMap[SessionToCreate.Session_key] = conn
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
