package jamf

import (
	"net/http"
	"time"
)

type Client struct {
	token, organization string

	// The Http Client that is used to make requests
	HttpClient       *http.Client
	HttpRetryTimeout time.Duration

	// Option to specify extra headers like User-Agent
	ExtraHeader map[string]string
}

// NewClient ... returns a new jamf.Client which can be used to access the API
func NewClient(token, organization string) *Client {
	return &Client{
		token:            token,
		organization:     organization,
		HttpClient:       http.DefaultClient,
		HttpRetryTimeout: 60 * time.Second,
		ExtraHeader:      make(map[string]string),
	}
}
