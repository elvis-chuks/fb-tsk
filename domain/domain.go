package domain

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Repository struct {
	CampaignRepo CampaignRepository
}

func ParseTime(timeStr string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, timeStr)
	return parsedTime
}

func HandleError(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(
		fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
}
