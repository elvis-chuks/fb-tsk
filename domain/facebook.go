package domain

import "time"

const (
	FacebookStatusActive   = "ACTIVE"
	FacebookStatusPaused   = "PAUSED"
	FacebookStatusDeleted  = "DELETED"
	FacebookStatusArchived = "ARCHIVED"
)

type FacebookCampaignResponse struct {
	Id                string      `json:"id"`
	AccountId         string      `json:"account_id"`
	Name              string      `json:"name"`
	Status            string      `json:"status"`
	Objective         string      `json:"objective"`
	DailyBudget       string      `json:"daily_budget"`
	LifetimeBudget    string      `json:"lifetime_budget"`
	StartTime         time.Time   `json:"start_time"`
	StopTime          time.Time   `json:"stop_time"`
	CreatedTime       time.Time   `json:"created_time"`
	UpdatedTime       time.Time   `json:"updated_time"`
	SpecialAdCategory string      `json:"special_ad_category"`
	SourceCampaignId  interface{} `json:"source_campaign_id"`
	ConfiguredStatus  string      `json:"configured_status"`
	EffectiveStatus   string      `json:"effective_status"`
	BuyingType        string      `json:"buying_type"`
	BudgetRemaining   string      `json:"budget_remaining"`
	PacingType        []string    `json:"pacing_type"`
}
