package business

import (
	"fmt"

	"github.com/allanfvc/finances/database"
	"github.com/allanfvc/finances/model"
	"github.com/allanfvc/finances/util"
)

//SaveInput creates a new input
func SaveInput(input *model.Input) {
	err := database.DBConn.Save(input).GetErrors()
	util.FormatErrors(err)
}

func ListInputs(name string) []model.Input {
	var inputs []model.Input
	var err []error
	if name != "" {
		query := fmt.Sprint("%%", name, "%%")
		err = database.DBConn.Where("name LIKE ? ", query).Find(&inputs).GetErrors()
	} else {
		err = database.DBConn.Find(&inputs).GetErrors()
	}
	util.FormatErrors(err)
	return inputs
}
