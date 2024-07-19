package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
	Tasks     []Task `gorm:"foreignKey:EmployeeID"`
}
