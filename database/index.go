package database

import (
	"fmt"

	"github.com/allanfvc/finances/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//DBConn database connection
var DBConn *gorm.DB

//InitDatabase opens the database connection
func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "./tmp/app.db")
	if err != nil {
		panic("failed to connect database")
	}
	DBConn.LogMode(true)
	DBConn.AutoMigrate(&model.User{})
	DBConn.AutoMigrate(&model.Input{})
	fmt.Println("Connection Opened to Database")
}

//CloseDatabase closes the database connection
func CloseDatabase() {
	if DBConn != nil {
		DBConn.Close()
	}
}
