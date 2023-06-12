package main

import (
	"github.com/ThanaponBunchot/demo-go-fiber/configs"
	"github.com/ThanaponBunchot/demo-go-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

type Contact struct {
	FirstName string `json:"firstName`
	LastName  string `json:"lastName`
	Tel       string `json:"tel"`
	Email     string `json:"email"`
	Detail    string `json:"detail"`
}

type Doctor struct {
	FirstName string `json:"firstName`
	LastName  string `json:"lastName`
	License   string `json:"license`
}

var doctor []Doctor

func main() {
	
	app := fiber.New()
	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)









	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value : " + c.Params("value"))
	})

	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("api path/" + c.Params("*"))
	})

	app.Static("/home", "./public")

	app.Post("/v1/contact", func(c *fiber.Ctx) error {
		ct := Contact{}
		err := c.BodyParser(&ct)
		if err != nil {
			return err

		} else {
			response := fiber.Map{
				"status": "Contact created successfully",
				"err":    err,
			}
			return c.JSON(response)
		}
	})

	app.Listen(":5000")
}
