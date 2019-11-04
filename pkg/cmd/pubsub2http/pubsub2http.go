package pubsub2http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/pubsub"
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
	log.Printf("Got Message from PubSub with ID %s", m.ID)

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Transferred Message with ID %s to URL: %s Got Response: %s", m.ID, p.PostURL, string(body))
}
