package controller

import (
	"strconv"

	"github.com/allanfvc/finances/business"
	"github.com/allanfvc/finances/model"
	"github.com/gofiber/fiber"
)

type userController struct {
}

func (c *userController) setupRoutes(app fiber.Router) {
	controller := "/users"
	app.Post(controller, save)
	app.Delete(controller+"/:id", delete)
}

func save(ctx *fiber.Ctx) {
	user := new(model.User)
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(503).Send(err)
		return
	}
	business.SaveUser(user)
	ctx.JSON(user)
}

func delete(ctx *fiber.Ctx) {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	business.DeleteUser(uint(id))
	ctx.Status(204)
}
