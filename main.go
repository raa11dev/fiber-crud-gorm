package main

import (
	"fiber-joglo-dev/database"
	"fiber-joglo-dev/database/migration"
	"fiber-joglo-dev/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial Database
	database.Init()
	// Initial Table Database
	migration.RunMigration()
	// Initial framework fiber
	app := fiber.New()
	// Initial Route
	routes.RouteInit(app)
	// Set log and address fiber
	log.Fatal(app.Listen(":3000"))
}
