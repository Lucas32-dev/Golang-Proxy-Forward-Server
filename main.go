package main

import (
	"fmt"
	"log"

	"github.com/Lucas32-dev/Golang-Proxy-Forward-Server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	// Get server config
	serverConfig := config.New()

	// Create app
	app := fiber.New(serverConfig.FiberConfig)

	// Log middleware
	app.Use(logger.New())

	// Forward requests
	for origin, destiny := range serverConfig.PathMap {
		app.Get(origin, proxy.Forward(destiny))
	}

	// Start server
	port := 3001
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
