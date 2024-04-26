package jamf

import (
	"errors"
	"time"
)

type responseOAuthToken struct {
	Token     *string `json:"access_token,omitempty"`
	ExpiresIn *int    `json:"expires_in,omitempty"`
}

type FormOptions struct {
	ClientId     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	GrantType    string `url:"grant_type"`
}

func (c *Client) CreateOAuthToken() (*responseOAuthToken, error) {
	if c.clientId == "" || c.clientSecret == "" {
		return nil, errors.New("Client ID and Client Secret are required to create an OAuth token")
	}

	var out *responseOAuthToken
	data := FormOptions{c.clientId, c.clientSecret, "client_credentials"}
	err := c.DoRequest("POST", uriOAuthToken, data, nil, &out)
	if err != nil {
		return nil, err
	}

	return out, err
}

type responseAuthToken struct {
	Token   *string    `json:"token,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
}

func (c *Client) CreateAccessToken() (*responseAuthToken, error) {
	if c.username == "" || c.password == "" {
		return nil, errors.New("username and password are required to create an Access Token")
	}

	var out *responseAuthToken
	err := c.DoRequest("POST", uriAuthToken, nil, nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
