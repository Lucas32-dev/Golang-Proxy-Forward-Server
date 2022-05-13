package config

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	PathMap     map[string]string
	FiberConfig fiber.Config
}

func New() Config {
	// Create path map
	pathMap := map[string]string{
		"/dog":             "https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/dog-puppy-on-garden-royalty-free-image-1586966191.jpg?crop=0.752xw:1.00xh;0.175xw,0&resize=640:*",
		"/jsonplaceholder": "https://jsonplaceholder.typicode.com/posts/1",
	}

	// Create fiber config
	fiberConfig := fiber.Config{
		GETOnly: true,
	}

	// Return config
	return Config{
		PathMap:     pathMap,
		FiberConfig: fiberConfig,
	}
}
