package memory

import (
	"mercurie/domain"
)

func NewMemoryDb() domain.Repository {

	var storage []domain.Campaign

	return domain.Repository{
		CampaignRepo: NewCampaignRepo(storage),
	}
}
