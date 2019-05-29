package nws

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	hClient   http.Client
	useragent string
}

// Create a new API client.
// A User Agent is required to identify your application. This string can be anything, and the more unique to your application the less likely it will be affected by a security event. If you include contact information (website or email), we can contact you if your string is associated to a security event.
func New(useragent string) *Client {
	// TODO
	return &Client{http.Client{}, useragent}
}

func (c *Client) Get(uri, contentType string) (*http.Response, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", contentType)
	req.Header.Set("User-Agent", c.useragent)

	resp, err := c.hClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		var svcError ServiceError
		dec := json.NewDecoder(resp.Body)
		err := dec.Decode(&svcError)
		if err != nil {
			return nil, fmt.Errorf("unable to decode API error: %v", err)
		}
		return nil, svcError
	}

	return resp, nil
}

// Perform a Get request with a ld+json content type.
func (c *Client) LDGet(uri string) (*http.Response, error) {
	return c.Get(uri, "application/ld+json")
}

// Perform a Get request with a geo+json content type.
func (c *Client) GeoGet(uri string) (*http.Response, error) {
	return c.Get(uri, "application/geo+json")
}

// Returned when there is a problem with the API.
type ServiceError struct {
	// A short, human-readable summary of the problem type
	Title string `json:"title"`
	// A URI reference (RFC3986) that identifies the problem type
	Type string `json:"type"`
	// The HTTP status code (RFC7231, Section 6) generated by the origin server for this occurrence of the problem
	Status int `json:"status"`
	// A human-readable explanation specific to this occurrence of the problem
	Detail string `json:"detail"`
	// A URI reference that identifies the specific occurrence of the problem
	Instance string `json:"instance"`
	// A unique identifier for the request, used for NWS debugging purposes. Please include this identifier with any correspondence to help us investigate your issue.
	CorrelationID string `json:"correlationId"`
}

func (s ServiceError) Error() string {
	return s.Detail
}
