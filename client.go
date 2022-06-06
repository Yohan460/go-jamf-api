package jamf

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	uriAuthToken = "/api/v1/auth/token"
)

// Client ... stores an object to talk with Jamf API
type Client struct {
	username, password, url string
	token                   *string
	tokenExpiration         *time.Time

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

func (c *Client) refreshAuthToken() error {
	if c.tokenExpiration != nil {
		if c.tokenExpiration.After(time.Now()) {
			return nil
		}
	}

	var out *responseAuthToken
	if err := c.doRequest("POST", uriAuthToken, nil, &out); err != nil {
		return err
	}
	c.token = out.Token
	c.tokenExpiration = out.Expires

	return nil
}
