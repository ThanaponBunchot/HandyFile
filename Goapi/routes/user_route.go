package routes

import (
	"github.com/ThanaponBunchot/demo-go-fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/user", controllers.CreateUser) //add this
}
