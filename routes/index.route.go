package routes

import (
	"fiber-joglo-dev/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerGetAll)
	r.Post("/user", handler.UserHandlerCreate)
	r.Get("/user/:nik", handler.UserHandlerGetById)
	r.Put("/user/:nik", handler.UserHandlerUpdate)
	r.Delete("/user/:nik", handler.UserHandlerDelete)
}
