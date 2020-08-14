package controller

import (
	"github.com/allanfvc/finances/business"
	"github.com/allanfvc/finances/model"
	"github.com/gofiber/fiber"
)

type inputController struct {
}

func (c *inputController) setupRoutes(app fiber.Router) {
	controller := "/inputs"
	app.Post(controller, saveInput)
	app.Get(controller, listInput)
}

func saveInput(ctx *fiber.Ctx) {
	input := new(model.Input)
	if err := ctx.BodyParser(input); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	business.SaveInput(input)
	ctx.JSON(input)
}

func listInput(ctx *fiber.Ctx) {
	name := ctx.Query("name")
	inputs := business.ListInputs(name)
	ctx.JSON(inputs)
}
