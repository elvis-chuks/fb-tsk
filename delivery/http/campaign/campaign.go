package campaign

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"mercurie/domain"
	"mercurie/pkg/queue"
	"strconv"
)

type handler struct {
	repo  domain.CampaignRepository
	queue *queue.CampaignQueue
}

func New(router fiber.Router, repo domain.CampaignRepository, campaignQueue *queue.CampaignQueue) {
	handler := handler{repo: repo, queue: campaignQueue}

	router.Post("/", handler.CreateCampaign)
	router.Get("/:id", handler.GetCampaign)
	router.Get("/", handler.GetAllCampaigns)
}

func (h handler) CreateCampaign(c *fiber.Ctx) error {

	var campaign domain.Campaign

	if err := json.Unmarshal(c.Body(), &campaign); err != nil {
		return domain.HandleError(c, err)
	}

	createdCampaign, err := h.repo.Create(campaign)

	if err != nil {
		return domain.HandleError(c, err)
	}

	err = h.queue.Enqueue(*createdCampaign)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"campaign": createdCampaign,
		},
	})
}

func (h handler) GetCampaign(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return domain.HandleError(c, err)
	}

	campaign, err := h.repo.Get(idInt)

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"campaign": campaign,
		},
	})
}

func (h handler) GetAllCampaigns(c *fiber.Ctx) error {

	campaigns, err := h.repo.GetAll()

	if err != nil {
		return domain.HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": fiber.Map{
			"campaigns": campaigns,
		},
	})
}
