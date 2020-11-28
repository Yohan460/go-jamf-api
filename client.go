package jamf

import (
	"fmt"
	"net/http"
	"time"
)

const (
	uriAuthToken = "/auth/tokens"
)

// Client ... stores an object to talk with Jamf API
type Client struct {
	username, password, url string
	token                   *string

	// The Http Client that is used to make requests
	HttpClient       *http.Client
	HttpRetryTimeout time.Duration

	// Option to specify extra headers like User-Agent
	ExtraHeader map[string]string
}

type ResponseAuthToken struct {
	Token   *string `json:"token,omitempty"`
	Expires *int    `json:"expires,omitempty"`
}

// NewClient ... returns a new jamf.Client which can be used to access the API
func NewClient(username, password, url string) (*Client, error) {
	c := &Client{
		username:         username,
		password:         password,
		url:              url,
		token:            nil,
		HttpClient:       http.DefaultClient,
		HttpRetryTimeout: 60 * time.Second,
		ExtraHeader:      make(map[string]string),
	}

	var out *ResponseAuthToken

	if err := c.doRequest("POST", uriAuthToken, nil, &out); err != nil {
		return nil, fmt.Errorf("cannot get a token: %v", err)
	}

	c.token = out.Token

	// initialize what don't need
	c.username = ""
	c.password = ""

	return c, nil
}
