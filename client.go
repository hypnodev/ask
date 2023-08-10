package ask

import (
	"context"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	ctx            context.Context
	httpClient     HttpClient
	baseUrl        string
	defaultHeaders http.Header
	verbose        bool
}

func NewClient(ctx context.Context) *Client {
	httpClient := &http.Client{}

	return &Client{
		ctx:            ctx,
		httpClient:     httpClient,
		defaultHeaders: http.Header{},
	}
}

func (client *Client) SetBaseUrl(url string) Client {
	client.baseUrl = url
	return *client
}

func (client *Client) AddDefaultHeader(key string, value string) Client {
	client.defaultHeaders.Set(key, value)
	return *client
}

func (client *Client) SetVerbose(flag bool) Client {
	client.verbose = flag
	return *client
}
