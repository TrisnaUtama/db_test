package server

import (
	"github.com/gofiber/fiber/v2"

	"base-1/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "base-1",
			AppName:      "base-1",
		}),

		db: database.New(),
	}

	return server
}
