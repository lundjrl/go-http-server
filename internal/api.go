package internal

import (
	"fmt"

	"github.com/lundjrl/go-http-server/internal/database"
	"github.com/lundjrl/go-http-server/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabaseConnection() {
	var err error

	database.DBConn, err = gorm.Open(sqlite.Open("app.db"))

	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection started")

	database.DBConn.AutoMigrate(&model.Cat{})

	fmt.Println("Database Migrated")
}
