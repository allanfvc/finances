package business

import (
	"strings"

	"github.com/allanfvc/finances/database"
	"github.com/allanfvc/finances/model"
)

func SaveUser(user *model.User) {
	err := database.DBConn.Save(user).GetErrors()
	if err != nil && len(err) > 0 {
		message := []string{}
		for _, element := range err {
			message = append(message, element.Error())
		}
		panic(strings.Join(message, ", "))
	}
}

func DeleteUser(id uint) {
	user := model.User{}
	user.ID = id
	err := database.DBConn.Delete(user).GetErrors()
	if err != nil && len(err) > 0 {
		message := []string{}
		for _, element := range err {
			message = append(message, element.Error())
		}
		panic(strings.Join(message, ", "))
	}
}
