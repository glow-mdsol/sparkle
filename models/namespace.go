package models

import "github.com/jinzhu/gorm"

type Namespace struct {
	gorm.Model
	Name    string	`json:"name",gorm:"type:varchar(50);unique_index"`
	URI		string	`json:"uri",gorm:"type:varchar(256);unique_index"`
}

