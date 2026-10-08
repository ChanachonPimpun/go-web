package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	// สร้าง instance ของ fiber
	app := fiber.New()

	// กำหนด route และ handler
	app.Get("/HelloGo", func(c fiber.Ctx) error {
		return c.SendString("Hello go TTTTT ....")
	})
	app.Get("/HiGo", func(c fiber.Ctx) error {
		return c.SendString("Hi go TT ....")
	})

	// start server
	if err := app.Listen(":4444"); err != nil {
		panic(err)
	}
}
