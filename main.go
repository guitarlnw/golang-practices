package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guitarlnw/golang-practices/controller"
)

func main() {
	app := fiber.New()

	uCtr := controller.NewUser()
	app.Get("/users/:id", uCtr.GetUser)

	app.Listen(":3000")
}
