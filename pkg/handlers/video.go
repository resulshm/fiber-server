package handlers

import (
	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) addVideo(c *fiber.Ctx) error {
	var input models.Video
	if input.Title = c.FormValue("title"); input.Title == "" {
		eMsg := "title of the video must be given"
		return newErrorResponse(c, 400, eMsg)
	}
	input.Description = c.FormValue("description")

	id, err := h.services.Video.AddVideo(input)
	if err != nil {
		eMsg := "Internal server error"
		return newErrorResponse(c, 500, eMsg)
	}
	logrus.Info("asdasdasd")
	return c.Status(200).JSON(map[string]int{
		"id": id,
	})
}
