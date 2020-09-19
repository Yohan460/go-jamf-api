package jamf

import (
	"net/http"
	"time"
)

type Client struct {
	token, organization string
	HttpClient          *http.Client
	HttpRetryTimeout    time.Duration
}

func NewClient(token, organization string) *Client {
	return &Client{
		token:            token,
		organization:     organization,
		HttpClient:       http.DefaultClient,
		HttpRetryTimeout: 60 * time.Second,
	}
}
