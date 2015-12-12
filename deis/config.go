package deis

import (
	"log"
	"net/url"

	deisClient "github.com/deis/deis/client/controller/client"
)

// Config holds API and APP keys to authenticate to Datadog.
type Config struct {
	ControllerURL string
	Token         string
	Username      string
}

// Client returns a new Deis client.
func (c *Config) Client() (*deisClient.Client, error) {
	controllerUrl, err := url.Parse(c.ControllerURL)
	if err != nil {
		return nil, err
	}

	client := &deisClient.Client{HTTPClient: deisClient.CreateHTTPClient(true), SSLVerify: true,
		ControllerURL: *controllerUrl, Token: c.Token, Username: c.Username,
		ResponseLimit: 100}

	log.Printf("[INFO] Deis Client configured ")

	return client, nil
}
