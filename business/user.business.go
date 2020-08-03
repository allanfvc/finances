package business

import (
	"fmt"

	"github.com/allanfvc/finances/util"

	"github.com/allanfvc/finances/database"
	"github.com/allanfvc/finances/model"
)

//SaveUser creates a new user
func SaveUser(user *model.User) {
	err := database.DBConn.Save(user).GetErrors()
	util.FormatErrors(err)
}

//DeleteUser remove a user from database
func DeleteUser(id uint) {
	user := model.User{}
	user.ID = id
	err := database.DBConn.Delete(user).GetErrors()
	util.FormatErrors(err)
}

//ListUsers list all users, can be filtered by user name
func ListUsers(name string) []model.User {
	var users []model.User
	var err []error
	if name != "" {
		query := fmt.Sprint("%%", name, "%%")
		err = database.DBConn.Where("name LIKE ? ", query).Find(&users).GetErrors()
	} else {
		err = database.DBConn.Find(&users).GetErrors()
	}
	util.FormatErrors(err)
	return users
}

//GetUser get user by id
func GetUser(id uint) model.User {
	user := model.User{}
	err := database.DBConn.First(&user, id).GetErrors()
	util.FormatErrors(err)
	return user
}
