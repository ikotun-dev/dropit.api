package models

import (
	"github.com/ikotun-dev/clipsync/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Session struct {
	gorm.Model
	Session_key string `gorm: ""json:"session_key"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Session{})
}

// function to create
func (ses *Session) CreateSession() *Session {
	db.NewRecord(ses)
	db.Create(&ses)
	return ses
}

// delete session
func DeleteSession(ID int64) Session {
	var session Session
	db.Where("ID?", ID).Delete(session)
	return session
}
