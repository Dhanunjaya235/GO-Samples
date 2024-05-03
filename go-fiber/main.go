package main

import (
	"fmt"
	"gofiber/apis/database"
	"gofiber/apis/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := database.Connect()
	defer database.DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	userRouter := app.Group("/users")

	userRouter.Get("/", handlers.GetAllUsers)
	userRouter.Post("/", handlers.AddNewUser)
	userRouter.Delete("/:id", handlers.DeleteUserByID)
	userRouter.Patch("/:id", handlers.UpdateUserById)
	userRouter.Get("/:id", handlers.GetUserByID)

	app.Listen(":4000")
}
