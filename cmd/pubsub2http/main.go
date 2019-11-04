package main

import (
	"log"
	"os"

	"github.com/300481/pubsub2http/pkg/cmd/pubsub2http"

	"github.com/urfave/cli"
)

var (
	app     = cli.NewApp()
	postUrl = os.Getenv("POST_URL")
)

func init() {
	env_vars := [...]string{
		"POST_URL",
		"GCP_CREDENTIALS_FILE",
		"GCP_TOPIC_NAME",
		"GCP_CREATE_TOPIC",
		"GCP_SUBSCRIPTION_NAME",
		"GCP_CREATE_SUBSCRIPTION",
		"GCP_PROJECT_ID"}

	for _, env_var := range env_vars {
		val, ok := os.LookupEnv(env_var)
		if !ok {
			log.Fatalf("Please set environment variable '%s'", env_var)
		} else {
			log.Printf("Environment Variable %s set to %s", env_var, val)
		}
	}
}

func info() {
	app.Name = "PubSub2HTTP"
	app.Usage = "A Bridge from Google PubSub to HTTP."
	app.Author = "Dennis Riemenschneider"
	app.Version = "0.1.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run in server mode",
			Action: func(c *cli.Context) {
				pubsub2http.New(postUrl).Serve()
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
