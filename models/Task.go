package models

import "gorm.io/gorm"

// Task model
type Task struct {
	gorm.Model
	Description   string
	StartDateTS   int32
	EndDateTS     int32
	EstimatedDays int32
	Status        string
	EmployeeID    uint
	ProjectID     uint
	Employee      Employee `gorm:"foreignKey:EmployeeID"`
	Project       Project  `gorm:"foreignKey:ProjectID"`
}
