package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/hng-task-1/controller"
)

func main() {
    app := fiber.New()

    app.Get("/api/classify-number",controller.GetFacts)

    log.Fatal(app.Listen(":3000"))
}