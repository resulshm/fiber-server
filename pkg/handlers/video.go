package handlers

import (
	"fmt"
	"strconv"

	"github.com/ResulShamuhammedov/fiber-server/pkg/models"
	"github.com/gofiber/fiber/v2"
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
		eMsg := fmt.Sprintf("Internal server error: %s", err.Error())
		return newErrorResponse(c, 500, eMsg)
	}

	return c.Status(200).JSON(map[string]int{
		"id": id,
	})
}

func (h *Handler) getVideoByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		eMsg := "id must be integer"
		return newErrorResponse(c, 400, eMsg)
	}

	video, err := h.services.Video.GetVideoByID(id)
	if err != nil {
		eMsg := fmt.Sprintf("Internal server error: %s", err.Error())
		return newErrorResponse(c, 500, eMsg)
	}
	if video.Id == 0 {
		eMsg := fmt.Sprintf("video with id=%d is not found", id)
		return newErrorResponse(c, 404, eMsg)
	}
	return c.Status(200).JSON(video)
}

func (h *Handler) getVideoByTitle(c *fiber.Ctx) error {
	var title string
	if title = c.Query("title"); title == "" {
		eMsg := "title must be given"
		return newErrorResponse(c, 400, eMsg)
	}

	video, err := h.services.Video.GetVideoByTitle(title)
	if err != nil {
		eMsg := fmt.Sprintf("Internal server error: %s", err.Error())
		return newErrorResponse(c, 500, eMsg)
	}
	if video.Id == 0 {
		eMsg := fmt.Sprintf("video with title=%s is not found", title)
		return newErrorResponse(c, 404, eMsg)
	}

	return c.Status(200).JSON(video)
}
