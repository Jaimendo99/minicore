package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name  string
	Tasks []Task `gorm:"foreignKey:ProjectID"`
}
