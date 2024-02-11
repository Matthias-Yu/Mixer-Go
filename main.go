package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()

	userSample := User{
		Name: "John",
		Age:  25,
	}

	userSampleStr := fmt.Sprintf("%v", userSample)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(userSampleStr)
	})


	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
