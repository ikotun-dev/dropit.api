package models

import (
	"github.com/ikotun-dev/clipsync/pkg/config"
	"github.com/jinzhu/gorm"
)

type Text struct {
	gorm.Model
	Content string `gorm:""json:"content"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Text{})
}

// function to create a clip
func (text *Text) CreateText() *Text {
	db.NewRecord(text)
	db.Create(&text)
	return text
}

// function to delte a clip
func DeleteText(ID int64) Text {
	var text Text
	db.Where("ID=?", ID).Delete(text)
	return text
}
