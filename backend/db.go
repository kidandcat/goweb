package backend

import (
	"goweb/backend/models"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	if app.IsClient {
		return
	}
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Don't forget to run migrations if you create a new model
	db.AutoMigrate(&models.Note{})
}
