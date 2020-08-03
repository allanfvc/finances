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
	app.Get(controller, list)
	app.Get(controller+"/:id", get)
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

func list(ctx *fiber.Ctx) {
	name := ctx.Query("name")
	users := business.ListUsers(name)
	ctx.JSON(users)
}

func get(ctx *fiber.Ctx) {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		panic(err)
	}
	user := business.GetUser(uint(id))
	ctx.JSON(user)
}
