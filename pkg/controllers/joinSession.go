package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JoinRequest struct {
	Sessionkey string `json:"session_key"`
}

// var db *gorm.DB
// function to join session
func JoinSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var joinReq JoinRequest
	err := json.NewDecoder(r.Body).Decode(&joinReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//vars := mux.Vars(r)
	session_key := joinReq.Sessionkey
	fmt.Println(session_key)
	if len(session_key) < 5 {
		res := map[string]string{"error": "session key should not be less than 7 characters"}
		errorMessage, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(errorMessage))

	} else {
		if ValidateSession(session_key) {

			//fmt.Println("session joined")
			w.WriteHeader(http.StatusOK)

		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("session key invalid")
		}

	}
}
