package database

import (
	"fmt"

	"github.com/allanfvc/finances/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "./tmp/app.db")
	if err != nil {
		panic("failed to connect database")
	}
	DBConn.LogMode(true)
	DBConn.AutoMigrate(&model.User{})
	fmt.Println("Connection Opened to Database")
}

func CloseDatabase() {
	if DBConn != nil {
		DBConn.Close()
	}
}
