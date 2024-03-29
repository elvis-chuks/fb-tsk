package domain

import "time"

type Campaign struct {
	Id              int                    `json:"id"`
	UserId          string                 `json:"user_id"`
	Budget          float64                `json:"budget"`
	Balance         float64                `json:"balance"`
	BudgetThreshold float64                `json:"budget_threshold"` // if my current campaign is <= this threshold, start pinging me
	Status          string                 `json:"status"`           // inactive, above_threshold, below_threshold
	MetaData        map[string]interface{} `json:"meta_data"`
	Retries         int                    `json:"retries"`
	NextRetry       time.Time              `json:"next_retry"`
}

type CampaignRepository interface {
	Create(campaign Campaign) (*Campaign, error)
	Get(id int) (*Campaign, error)
	GetAll() ([]Campaign, error)
	GetLastCampaignId() (int, error)
	PopulateCampaigns() error
}
