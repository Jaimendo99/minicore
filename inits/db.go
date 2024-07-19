package inits

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"minicore/models"
)

var DB *gorm.DB

func NewConnection(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}
	return db, nil
}

func InitDB() {
	db, err := NewConnection("test.db")
	if err != nil {
		fmt.Println("error opening db" + err.Error())
	}
	DB = db
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Task{}, &models.Employee{}, &models.Project{})
	if err != nil {
		return fmt.Errorf("Error migrating database: %v", err)
	}
	return nil
}
