package jamf

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	uriAuthToken  = "/api/v1/auth/token"
	uriOAuthToken = "/api/oauth/token"
)

// Client ... stores an object to talk with Jamf API
type Client struct {
	username, password, url, clientId, clientSecret string
	token                                           *string
	tokenExpiration                                 *time.Time

	// The Http Client that is used to make requests
	HttpClient       *http.Client
	HttpRetryTimeout time.Duration

	// Option to specify extra headers like User-Agent
	ExtraHeader map[string]string
}

type responseAuthToken struct {
	Token   *string    `json:"token,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
}

type responseOAuthToken struct {
	Token     *string `json:"access_token,omitempty"`
	ExpiresIn *int    `json:"expires_in,omitempty"`
}

type FormOptions struct {
	ClientId     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	GrantType    string `url:"grant_type"`
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

	if err := c.refreshAuthToken(); err != nil {
		return c, errors.Wrap(err, "Error getting bearer auth token")
	}

	return c, nil
}

// NewOAuthClient ... returns a new jamf.Client which can be used to access the API using the new bearer tokens
func NewOAuthClient(clientId, clientSecret, url string) (*Client, error) {
	c := &Client{
		clientId:         clientId,
		clientSecret:     clientSecret,
		url:              url,
		token:            nil,
		HttpClient:       http.DefaultClient,
		HttpRetryTimeout: 60 * time.Second,
		ExtraHeader:      make(map[string]string),
	}

	if err := c.refreshAuthToken(); err != nil {
		return c, errors.Wrap(err, "Error getting bearer auth token")
	}

	return c, nil
}

func (c *Client) refreshAuthToken() error {
	if c.tokenExpiration != nil {
		if c.tokenExpiration.After(time.Now()) {
			return nil
		}
	}

	c.token = nil

	var out *responseAuthToken
	if c.clientId != "" && c.clientSecret != "" {
		var out *responseOAuthToken
		data := FormOptions{c.clientId, c.clientSecret, "client_credentials"}
		if err := c.DoRequest("POST", uriOAuthToken, data, nil, &out); err != nil {
			return err
		}

		c.token = out.Token
		expiration := time.Now().Add(time.Duration(*out.ExpiresIn) * time.Second)
		c.tokenExpiration = &expiration
	} else {
		if err := c.DoRequest("POST", uriAuthToken, nil, nil, &out); err != nil {
			return err
		}

		c.token = out.Token
		c.tokenExpiration = out.Expires
	}

	return nil
}
