package controllers

import (
	"github.com/ikotun-dev/clipsync/pkg/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CheckSession(SessionKey string) bool {
	var session models.Session

	if db.First(&session, "session_key=:?", SessionKey).RecordNotFound() {
		return false
	} else {
		return true
	}
}
