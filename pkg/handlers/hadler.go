package handlers

import (
	"github.com/ResulShamuhammedov/fiber-server/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	app.Post("/video", h.addVideo)

	return app
}
