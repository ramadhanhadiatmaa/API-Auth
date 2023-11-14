package routes

import (
	"apiauth/controllers"
	"apiauth/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	auth := r.Group("/api")

	auth.Get("/:username", middlewares.Auth, controllers.Show)
	auth.Post("/", middlewares.Auth, controllers.Create)
	auth.Put("/:username", middlewares.Auth, controllers.Update)
	auth.Delete("/:username", middlewares.Auth, controllers.Delete)
}