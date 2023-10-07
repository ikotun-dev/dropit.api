package controllers

import (
	"github.com/ikotun-dev/clipsync/pkg/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CheckSession(SessionKey string) bool {
	var session *models.Session

	if err := db.Where("session_key = ?", SessionKey).First(&session).Error; err != nil {
		return false
	} else {
		return true
	}
}

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
