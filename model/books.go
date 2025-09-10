package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	ID     uint `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}	
