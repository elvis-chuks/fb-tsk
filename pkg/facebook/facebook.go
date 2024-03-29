package facebook

import (
	"mercurie/domain"
	"net/http"
)

func GetCampaignData(id int) (*domain.FacebookCampaignResponse, error) {
	// mock http request to facebooks api

	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	_, err = client.Do(req)

	if err != nil {
		return nil, err
	}

	return &domain.FacebookCampaignResponse{
		Id:                "1234567890",
		AccountId:         "987654321",
		Name:              "My Campaign",
		Status:            "ACTIVE",
		Objective:         "LINK_CLICKS",
		DailyBudget:       "100.00",
		LifetimeBudget:    "1000.00",
		StartTime:         domain.ParseTime("2024-03-01T00:00:00+00:00"),
		StopTime:          domain.ParseTime("2024-03-31T00:00:00+00:00"),
		CreatedTime:       domain.ParseTime("2024-03-01T08:00:00+00:00"),
		UpdatedTime:       domain.ParseTime("2024-03-01T08:00:00+00:00"),
		SpecialAdCategory: "NONE",
		SourceCampaignId:  nil,
		ConfiguredStatus:  "ACTIVE",
		EffectiveStatus:   "ACTIVE",
		BuyingType:        "AUCTION",
		BudgetRemaining:   "500.00",
		PacingType:        []string{"standard"},
	}, nil
}
