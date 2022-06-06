package jamf

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/cenkalti/backoff"
)

const (
	duplicateNameErr string = "Duplicate Name"
)

type Error interface {
	Error() string
	StatusCode() int
	URI() string
	Body() string
}
type errorInfo struct {
	statusCode int
	uri        string
	body       string
}

func newError(statusCode int, uri, body string) Error {
	var e = new(errorInfo)
	e.statusCode = statusCode
	e.uri = uri
	e.body = body
	return e
}

func (e *errorInfo) Error() string {
	return fmt.Sprintf("API Error: %v URI: %s Body: %s", e.statusCode, e.uri, e.body)
}

func (e *errorInfo) StatusCode() int {
	return e.statusCode
}

func (e *errorInfo) URI() string {
	return e.uri
}

func (e *errorInfo) Body() string {
	return e.body
}

// doJsonRequest ... A method to send a request to the jamf api
func (c *Client) doRequest(method, api string, reqbody interface{}, params *url.Values, out interface{}) error {
	req, err := c.createRequest(method, api, params, reqbody)
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
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		re := regexp.MustCompile(`\r?\n`)
		out := re.ReplaceAllString(string(body), " ")
		return newError(resp.StatusCode, api, out)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Making sure we have someplace to put the response
	if out != nil {

		// If we got no body, by default let's just make an empty JSON dict.
		if len(body) == 0 {
			body = []byte{'{', '}'}
		}

		if strings.Contains(api, "JSSResource") {
			err = xml.Unmarshal(body, out)
		} else {
			err = json.Unmarshal(body, out)
		}
	}

	return err
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
	return fmt.Sprintf("https://%s%s", c.url, api)
}

// createRequest ...　Generate a http request for api.
func (c *Client) createRequest(method, api string, params *url.Values, reqbody interface{}) (*http.Request, error) {
	var bodyReader io.Reader

	// Convert the request body to the appropriate type
	if method != "GET" && reqbody != nil {

		if strings.Contains(api, "JSSResource") {
			b, err := xml.Marshal(reqbody)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewReader(b)
		} else {
			b, err := json.Marshal(reqbody)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequest(method, c.uriForAPI(api), bodyReader)
	if err != nil {
		return nil, err
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	// Ensuring valid token
	if c.token != nil {
		err = c.refreshAuthToken()
		if err != nil {
			return req, err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *c.token))
	} else {
		// Handle inital token setup in client setup
		req.SetBasicAuth(c.username, c.password)
	}

	// Set the necessary headers
	if strings.Contains(api, "JSSResource") || c.token == nil {
		req.Header.Add("Content-Type", "application/xml")
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	for k, v := range c.ExtraHeader {
		req.Header.Add(k, v)
	}

	return req, err
}
