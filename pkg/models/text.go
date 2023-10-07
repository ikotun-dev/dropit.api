package models

import (
	"github.com/ikotun-dev/clipsync/pkg/config"
	"github.com/jinzhu/gorm"
)

type Text struct {
	gorm.Model
	Content   string  `gorm:""json:"content"`
	SessionID Session `gorm:"foreignkey:SessionID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Text{})
}

// function to create a clip
func (text *Text) CreateText() error {
	db.NewRecord(text)
	if err := db.Create(&text).Error; err != nil {
		return err
	}
	return nil
}

// function to delte a clip
func DeleteText(ID int64) error {
	var text Text
	return db.Where("ID=?", ID).Delete(text).Error

}
