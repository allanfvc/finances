package controller

import "github.com/gofiber/fiber"

type appController interface {
	setupRoutes(app fiber.Router)
}

//RegisterRoutes register all app routes
func RegisterRoutes(app fiber.Router) {
	var routes = getAllRoutes()
	for _, route := range routes {
		route.setupRoutes(app)
	}
}

func getAllRoutes() []appController {
	routes := []appController{}
	routes = append(routes, new(userController))
	routes = append(routes, new(inputController))
	return routes
}
