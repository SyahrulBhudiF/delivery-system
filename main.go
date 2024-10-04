package main

import (
	"delivery-system/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Rute untuk membuat pengiriman
	app.Post("/api/deliveries", handlers.CreateDelivery)

	// Rute untuk mendapatkan pengiriman berdasarkan ID
	app.Get("/api/deliveries/:id", handlers.GetDelivery)

	// Menjalankan server pada port 3000
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
