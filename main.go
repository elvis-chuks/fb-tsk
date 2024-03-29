package main

import (
	"flag"
	"fmt"
	"log"
	httpDelivery "mercurie/delivery/http"
	"mercurie/pkg/queue"
	"mercurie/repository/memory"
	"os"
	"time"
)

// Design a web service to connect to Facebook API to pull campaign data and send notifications to user when campaign budget is below a certain threshold

func main() {

	mode := flag.String("mode", "client", "Mode of execution")
	flag.Parse()

	repo := memory.NewMemoryDb()

	switch *mode {
	case "server":
		campaignQueue := queue.CreateQueue(nil, time.Second*1)
		httpConfig := httpDelivery.Config{CampaignRepo: repo.CampaignRepo, CampaignQueue: &campaignQueue}

		app := httpDelivery.SetupRouter(httpConfig)

		port := os.Getenv("PORT")

		if port == "" {
			port = "5001"
		}

		done := make(chan bool)

		go campaignQueue.Worker(done, true)

		addr := flag.String("addr", fmt.Sprintf(":%s", port), "http service address")
		flag.Parse()
		log.Fatal(app.Listen(*addr))
	default:
		// populate queue, pull data from hypothetical source

		repo := memory.NewMemoryDb().CampaignRepo

		err := repo.PopulateCampaigns()

		if err != nil {
			panic(err)
		}

		campaigns, err := repo.GetAll()

		if err != nil {
			panic(err)
		}

		campaignQueue := queue.CreateQueue(campaigns, time.Second*1)

		done := make(chan bool)

		go campaignQueue.Worker(done, false) // run queue worker

		<-done
	}
}
