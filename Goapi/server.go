// package main

// import (
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// )

// type Contact struct {
// 	FirstName    string `form:"firstName"`
// 	LastName     string `form:"lastName`
// 	Email        string `form:"email"`
// 	MobileNumber string `form:"mobileNumber"`
// 	Detail       string `form:"detail"`
// }

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})

// 	app.Post("/v1/contact", func(c *fiber.Ctx) error {
// 		ct := Contact{}
// 		if err := c.BodyParser(&ct); err != nil {
// 			return err
// 		}
// 		log.Println(ct.FirstName)
// 		log.Println(ct.LastName)
// 		return c.JSON(ct)
// 	})

// 	app.Listen(":3000")
// }

package main

import "github.com/gofiber/fiber/v2"

type Contact struct {
	FirstName string `form:"firstName`
	LastName  string `form:"lastName`
	Tel       string `form:"tel"`
	Email     string `form:"email"`
	Detail    string `from:"detail"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

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

	app.Listen(":3000")
}
