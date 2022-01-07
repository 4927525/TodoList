package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Uid uint `gorm:"type:int;not null"`
	Title string `gorm:"type:varchar(255);not null"`
	Status uint `gorm:"type:int;not null;default:0"`
	Content string `gorm:"type:text"`
	StartTime int64
	EndTime int64
}
