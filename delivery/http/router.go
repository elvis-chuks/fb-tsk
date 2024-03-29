package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mercurie/delivery/http/campaign"
	"mercurie/domain"
	"mercurie/pkg/queue"
)

type Config struct {
	CampaignRepo  domain.CampaignRepository
	CampaignQueue *queue.CampaignQueue
}

func SetupRouter(config Config) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	v1 := app.Group("/api/v1")

	campaign.New(v1, config.CampaignRepo, config.CampaignQueue)

	return app
}
