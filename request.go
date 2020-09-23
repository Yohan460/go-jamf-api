package jamf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/cenkalti/backoff"
)

// doJsonRequest ... A method to send a request to the jamf api
func (c *Client) doRequest(method, api string, reqbody, out interface{}) error {
	if out == nil {
		return nil
	}

	req, err := c.createRequest(method, api, reqbody)
	if err != nil {
		return err
	}

	// request
	var resp *http.Response
	if method == "POST" || method == "PUT" {
		resp, err = c.HttpClient.Do(req)
	} else {
		resp, err = c.doRequestWithRetries(req)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("api error %s: it's possible that the Token/Organization is just wrong", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// If we got no body, by default let's just make an empty JSON dict.
	if len(body) == 0 {
		body = []byte{'{', '}'}
	}

	return json.Unmarshal(body, &out)
}

// doRequestWithRetries ... GET/DELETE depends on the jamf server, the retry process is largely
// Use backoff to extend the wait interval for retrying exponentially.
func (c *Client) doRequestWithRetries(req *http.Request) (*http.Response, error) {
	var (
		err  error
		resp *http.Response
		bo   = backoff.NewExponentialBackOff()
		body []byte
	)

	bo.MaxElapsedTime = c.HttpRetryTimeout

	if req.Body != nil {
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return resp, err
		}
	}

	boReq := func() error {
		if body != nil {
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		resp, err = c.HttpClient.Do(req)
		if err != nil {
			return err
		}

		// 2xx are done. 4xx are not retry
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return nil
		}

		return fmt.Errorf("received http status code %d", resp.StatusCode)
	}

	err = backoff.Retry(boReq, bo)
	return resp, err
}

// uriForApi ... Generate uri for api
func (c *Client) uriForAPI(api string) string {
	return fmt.Sprintf("https://%s.jamfcloud.com/uapi%s", c.organization, api)
}

// createRequest ...　Generate a http request for api.
func (c *Client) createRequest(method, api string, reqbody interface{}) (*http.Request, error) {
	var bodyReader io.Reader

	// Convert the request body to json
	if method != "GET" && reqbody != nil {
		b, err := json.Marshal(reqbody)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, c.uriForAPI(api), bodyReader)
	if err != nil {
		return nil, err
	}

	// Set the necessary headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	return req, nil
}
