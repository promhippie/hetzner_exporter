package hetzner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/promhippie/hetzner_exporter/pkg/version"
)

var (
	// UserAgent defines the used user ganet for request.
	UserAgent = fmt.Sprintf("hetzner_export/%s", version.String)
)

// Client is a client for the DockerHub API.
type Client struct {
	httpClient *http.Client
	httpDumper Dumper
	username   string
	password   string

	Server     ServerClient
	SSHKey     SSHKeyClient
	Storagebox StorageboxClient
}

// A ClientOption is used to configure a Client.
type ClientOption func(*Client)

// WithHTTPClient configures a Client to use the specified HTTP client.
func WithHTTPClient(value *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = value
	}
}

// WithHTTPDumper configures a Client to use the specified debug dumper.
func WithHTTPDumper(value Dumper) ClientOption {
	return func(client *Client) {
		client.httpDumper = value
	}
}

// WithUsername configures a Client to use the specified username for authentication.
func WithUsername(value string) ClientOption {
	return func(client *Client) {
		client.username = value
	}
}

// WithPassword configures a Client to use the specified password for authentication.
func WithPassword(value string) ClientOption {
	return func(client *Client) {
		client.password = value
	}
}

// NewClient creates a new client.
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{},
	}

	for _, option := range options {
		option(client)
	}

	client.Server = ServerClient{client: client}
	client.SSHKey = SSHKeyClient{client: client}
	client.Storagebox = StorageboxClient{client: client}

	return client
}

// NewRequest creates an HTTP request against the DockerHub.
func (c *Client) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)

	if c.username != "" && c.password != "" {
		req.SetBasicAuth(
			c.username,
			c.password,
		)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req.WithContext(ctx), nil
}

// Do performs an HTTP request against the DockerHub.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	if c.httpDumper != nil {
		c.httpDumper.DumpRequest(req)
	}

	res, err := c.httpClient.Do(req)

	if res != nil {
		defer func() { _ = res.Body.Close() }()
	}

	if err != nil {
		return nil, err
	}

	if c.httpDumper != nil {
		c.httpDumper.DumpResponse(res)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return &Response{Response: res}, err
	}

	res.Body = io.NopCloser(bytes.NewReader(body))

	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		return &Response{Response: res}, errors.New(http.StatusText(res.StatusCode))
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, bytes.NewReader(body))
		} else {
			err = json.Unmarshal(body, v)
		}
	}

	return &Response{Response: res}, err
}

// Response simply wraps the standard response type.
type Response struct {
	*http.Response
}
