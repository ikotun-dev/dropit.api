package models

import (
	"log"

	"github.com/ikotun-dev/clipsync/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Session struct {
	gorm.Model
	Session_key string `gorm:"column:session_key" json:"session_key"`
	Text        []Text `gorm:"foreign:SessionID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Session{})
}

func (ses *Session) CreateSession() error {
	log.Println("Creating session...")

	db.NewRecord(ses)
	if err := db.Create(&ses).Error; err != nil {
		log.Println(err)
		return err
	}

	log.Println("Session created successfully.")
	return nil
}

func GetSessionByKey(session_key string) (*Session, *gorm.DB) {
	var session Session
	db := db.Where("session_key=?", session_key).Find(&session)
	return &session, db
}

// delete session
func DeleteSession(ID int64) Session {
	var session Session
	db.Where("ID?", ID).Delete(session)
	return session
}
