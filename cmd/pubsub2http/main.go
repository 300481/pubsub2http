package main

import (
	"log"
	"os"

	"github.com/300481/mq"
	"github.com/300481/pubsub2http/pkg/cmd/pubsub2http"
)

// newGCP creates new GCP PubSub consumer struct
func newGCP() *mq.GCP {
	return &mq.GCP{
		CredentialsFile:    os.Getenv("GCP_CREDENTIALS_FILE"),
		TopicName:          os.Getenv("GCP_TOPIC_NAME"),
		CreateTopic:        os.Getenv("GCP_CREATE_TOPIC") == "TRUE",
		SubscriptionName:   os.Getenv("GCP_SUBSCRIPTION_NAME"),
		CreateSubscription: os.Getenv("GCP_CREATE_SUBSCRIPTION") == "TRUE",
		ProjectID:          os.Getenv("GCP_PROJECT_ID"),
	}
}

func main() {
	log.Println("Start PubSub2HTTP")

	// create new configuration from environment
	p := pubsub2http.New(os.Getenv("POST_URL"))

	// create new message queue configuration from environment
	mq := newGCP()

	// subscribe to the message queue
	err := mq.Subscribe(p.HandleMessage)
	if err != nil {
		log.Fatalln(err)
	}
}
