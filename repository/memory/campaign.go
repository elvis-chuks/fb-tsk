package memory

import (
	"errors"
	"math/rand"
	"mercurie/domain"
)

type campaignHandler struct {
	repo []domain.Campaign
}

func (c *campaignHandler) PopulateCampaigns() error {
	for i := 0; i <= 50; i++ {
		lastId, err := c.GetLastCampaignId()

		if err != nil {
			return err
		}
		_, err = c.Create(domain.Campaign{
			Id:              lastId + 1,
			UserId:          "12345",
			Budget:          float64(rand.Intn(700-500) + 400),
			BudgetThreshold: float64(rand.Intn(700-500) + 400),
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *campaignHandler) GetLastCampaignId() (int, error) {

	if len(c.repo) == 0 {
		return 0, nil
	}

	return c.repo[len(c.repo)-1].Id, nil
}

func (c *campaignHandler) Create(campaign domain.Campaign) (*domain.Campaign, error) {
	lastId, err := c.GetLastCampaignId()

	if err != nil {
		return nil, err
	}

	campaign.Id = lastId + 1

	//if campaign.Budget <= 0 {
	//	return nil, errors.New("campaign budget must be greater than zero")
	//}
	//
	//if campaign.BudgetThreshold > campaign.Budget {
	//	return nil, errors.New("budget threshold must be less than budget")
	//}

	c.repo = append(c.repo, campaign)

	return &campaign, nil
}

func (c *campaignHandler) Get(id int) (*domain.Campaign, error) {

	for _, i := range c.repo {
		if i.Id == id {

			return &i, nil
		}
	}

	return nil, errors.New("campaign with that id does not exist")
}

func (c *campaignHandler) GetAll() ([]domain.Campaign, error) {
	return c.repo, nil
}

func NewCampaignRepo(repo []domain.Campaign) domain.CampaignRepository {
	return &campaignHandler{repo: repo}
}
