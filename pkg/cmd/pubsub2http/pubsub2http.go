package pubsub2http

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/300481/mq"
)

// PubSub2HTTP the struct for handle functions
type PubSub2HTTP struct {
	PostURL string
}

// pubSubMessage contains the type of webhook and its payload
type pubSubMessage struct {
	Method string
	Header http.Header
	Body   []byte
}

// New creates new PubSub2HTTP struct
func New(postUrl string) *PubSub2HTTP {
	return &PubSub2HTTP{
		PostURL: postUrl,
	}
}

// HandleMessage handles incoming messages
func (p *PubSub2HTTP) HandleMessage(ctx context.Context, m *pubsub.Message) {
	defer m.Ack()

	var message pubSubMessage
	err := json.Unmarshal(m.Data, &message)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Got Message from PubSub [ID: %s]", m.ID)

	request, err := http.NewRequest(message.Method, p.PostURL, bytes.NewReader(message.Body))
	if err != nil {
		log.Println(err)
		return
	}
	request.Header = message.Header

	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Transferred Message [ID: %s] [URL: %s]", m.ID, p.PostURL)
}

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

// Server starts the server
func (p *PubSub2HTTP) Serve() {
	log.Println("Start PubSub2HTTP")

	// create new message queue configuration from environment
	mq := newGCP()

	// subscribe to the message queue
	err := mq.Subscribe(p.HandleMessage)
	if err != nil {
		log.Fatalln(err)
	}
}
