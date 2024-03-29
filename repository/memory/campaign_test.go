package memory

import (
	"fmt"
	"mercurie/domain"
	"testing"
)

var repo = NewMemoryDb().CampaignRepo

func TestCampaignHandler_Create(t *testing.T) {

	campaign, err := repo.Create(domain.Campaign{
		UserId:          "1234",
		Budget:          1000,
		BudgetThreshold: 400,
	})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%+v \n", campaign)
}
