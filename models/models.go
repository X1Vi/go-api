package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// User model
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"` // Username must be unique and non-null
	ToDos    []ToDo // Define the relationship between User and ToDo
}

// ToDo model
type ToDo struct {
	gorm.Model
	ToDo   string `gorm:"not null"`      // ToDo field should be non-null
	IsDone bool   `gorm:"default:false"` // IsDone defaults to false
	UserID uint   // Foreign key for User
	User   User   `gorm:"foreignKey:UserID"` // Define relationship
}

// MigrateModels initializes the database and runs schema migrations
func MigrateModels() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // Open the database connection and assign to the global db
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema for all models
	db.AutoMigrate(&User{}, &ToDo{})
}

// GetDb returns the initialized database connection
func GetDb() *gorm.DB {
	return db
}
