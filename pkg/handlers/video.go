package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) addVideo(c *fiber.Ctx) error {

	return c.SendString("Hello fiber")
}
