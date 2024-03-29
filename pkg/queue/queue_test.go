package queue

import (
	"fmt"
	"mercurie/domain"
	"mercurie/repository/memory"
	"testing"
)

var repo = memory.NewMemoryDb().CampaignRepo

var campaign, _ = repo.Create(domain.Campaign{
	UserId:          "1234",
	Budget:          1000,
	BudgetThreshold: 400,
})

var queue = CampaignQueue{}

func TestCampaignQueue_Enqueue_Dequeue(t *testing.T) {

	err := queue.Enqueue(*campaign)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(queue)

	dequeued, err := queue.Dequeue()

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(dequeued)

	fmt.Println(queue)
}
