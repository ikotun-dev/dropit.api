package controllers

import (
	"github.com/ikotun-dev/clipsync/pkg/models"
)

func ValidateSession(session_key string) bool {
	// Get the session by session_key
	session, db := models.GetSessionByKey(session_key)

	// Check if the session is valid
	if db.Error != nil || session.ID == 0 {
		// Session is not valid
		return false
	}

	// Session is valid
	return true
}
