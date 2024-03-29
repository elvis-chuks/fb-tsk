package queue

import (
	"errors"
	"fmt"
	"mercurie/domain"
	"mercurie/pkg/facebook"
	"mercurie/pkg/olugbe_ilu"
	"strconv"
	"time"
)

const MaxRetries = 3

type CampaignQueue struct {
	Campaigns []domain.Campaign
	Ticker    *time.Ticker
}

func CreateQueue(campaigns []domain.Campaign, frequency time.Duration) CampaignQueue {
	return CampaignQueue{Campaigns: campaigns, Ticker: time.NewTicker(frequency)}
}

func (q *CampaignQueue) Enqueue(campaign domain.Campaign) error {
	q.Campaigns = append(q.Campaigns, campaign)
	return nil
}

func (q *CampaignQueue) Length() int {
	return len(q.Campaigns)
}

func (q *CampaignQueue) Dequeue() (*domain.Campaign, error) {
	if len(q.Campaigns) == 0 {
		return nil, errors.New("cannot dequeue empty queue")
	}

	lastCampaign := q.Campaigns[0]

	q.Campaigns = q.Campaigns[1:]

	return &lastCampaign, nil
}

func (q *CampaignQueue) Worker(done chan<- bool, isServer bool) {
	for {

		select {
		case _ = <-q.Ticker.C: // run everytime it ticks, so you don't clog the system with unnecessary operations
			if !isServer {
				if q.Length() == 0 {
					done <- true
				}
			}

			if isServer {
				if q.Length() == 0 {
					continue
				}
			}

			campaign, err := q.Dequeue()

			if err != nil {
				// an error occurred
				// maybe send it back to the queue and try again later
				q.HandleWorkerError(campaign, isServer)
				continue
			}

			if isServer {
				if time.Now().Before(campaign.NextRetry) {
					q.Enqueue(*campaign)
					continue
				}
			}

			if campaign.Retries > MaxRetries {
				// take out from queue
				continue
			}

			fmt.Println(campaign)

			// check campaign status from facebook

			facebookResponse, err := facebook.GetCampaignData(campaign.Id)

			if err != nil {
				q.HandleWorkerError(campaign, isServer)
				continue
			}

			// check campaign status

			switch facebookResponse.Status {
			case domain.FacebookStatusActive:
				// remind user

				value, err := strconv.ParseFloat(facebookResponse.BudgetRemaining, 64)

				if err != nil {
					q.HandleWorkerError(campaign, isServer)
					continue
				}

				fmt.Println(value, campaign.BudgetThreshold)

				if value < campaign.BudgetThreshold {
					fmt.Println("go and top up your budget o")
					go olugbe_ilu.NotifyUserAboutCampaignBudget(*campaign)
					q.HandleWorkerError(campaign, isServer)
				} else {
					if isServer {
						q.HandleWorkerSuccess(campaign, isServer)
					}

				}

			case domain.FacebookStatusDeleted, domain.FacebookStatusArchived:
				// do nothing
			case domain.FacebookStatusPaused:
				// notify user to resume budget
				go olugbe_ilu.NotifyUserAboutResumingCampaign(*campaign)
			default:
				// do nothing
			}
		}

	}
}

func (q *CampaignQueue) HandleWorkerError(campaign *domain.Campaign, isServer bool) {
	if campaign != nil {
		campaign.Retries += 1
		if isServer {
			campaign.NextRetry = time.Now().Add(time.Minute * 1)
		}
		err := q.Enqueue(*campaign)
		if err != nil {
			return
		}
	}
}

func (q *CampaignQueue) HandleWorkerSuccess(campaign *domain.Campaign, isServer bool) {
	if campaign != nil {
		campaign.Retries = 0
		if isServer {
			campaign.NextRetry = time.Now().Add(time.Minute * 1)
		}
		err := q.Enqueue(*campaign)
		if err != nil {
			return
		}
	}
}
