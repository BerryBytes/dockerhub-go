package dockerhub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	defaultUserAgent       = "dockerhub-go/v1"
	defaultAPIBaseURL      = "https://hub.docker.com"
	defaultAPIBaseEndpoint = "/v2"
)

// A Client manages communication with the Dockerhub API.
type Client struct {
	httpClient *http.Client
	BaseURL    *url.URL
	UserAgent  string

	authToken string

	common service

	Auth         *AuthService
	Repositories *RepositoriesService
}

// NewClient returns a new Dockerhub client. If an httpClient is not
// provided, a new http.Client will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultAPIBaseURL)

	c := &Client{
		httpClient: httpClient,
		UserAgent:  defaultUserAgent,
		BaseURL:    baseURL,
	}
	c.common.client = c
	c.Auth = (*AuthService)(&c.common)
	c.Repositories = (*RepositoriesService)(&c.common)
	return c
}

type service struct {
	client *Client
}

// SetAuthToken sets the Authorization token on the client to be sent with
// API requests.
func (c *Client) SetAuthToken(token string) {
	c.authToken = token
}

// Do sends an API request and returns the API response. The API response is JSON
// decoded and stored in the value pointed to by v.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return err
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil && err != io.EOF {
		return err
	}
	return nil
}

// NewRequest creates an API request. The given URL is relative to the Client's
// BaseURL.
func (c *Client) NewRequest(method, url string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(defaultAPIBaseEndpoint + url)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if len(c.authToken) != 0 {
		req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.authToken))
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("User-Agent", c.UserAgent)
	debug, _ := httputil.DumpRequestOut(req, true)
	fmt.Println(string(debug))
	return req, nil
}
