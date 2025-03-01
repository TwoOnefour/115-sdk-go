package _115

import (
	"encoding/json"
	"net/http"

	"resty.dev/v3"
)

type Client struct {
	client            *resty.Client
	jsonMarshalFunc   func(v interface{}) ([]byte, error)
	jsonUnmarshalFunc func(data []byte, v interface{}) error

	accessToken    string
	refreshToken   string
	onRefreshToken func(string, string)
}

func New(opts ...Option) *Client {
	c := &Client{
		client:            resty.New(),
		jsonMarshalFunc:   json.Marshal,
		jsonUnmarshalFunc: json.Unmarshal,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func Default() *Client {
	return New()
}

func (w *Client) SetHttpClient(httpClient *http.Client) *Client {
	w.client = resty.NewWithClient(httpClient)
	return w
}

func (w *Client) SetUserAgent(userAgent string) *Client {
	w.client.SetHeader("User-Agent", userAgent)
	return w
}

func (w *Client) SetDebug(d bool) *Client {
	w.client.SetDebug(d)
	return w
}

func (w *Client) EnableTrace() *Client {
	w.client.EnableTrace()
	return w
}

func (w *Client) SetProxy(proxy string) *Client {
	w.client.SetProxy(proxy)
	return w
}

func (w *Client) SetJsonMarshalFunc(f func(v interface{}) ([]byte, error)) {
	w.jsonMarshalFunc = f
}

func (w *Client) SetJsonUnmarshalFunc(f func(data []byte, v interface{}) error) {
	w.jsonUnmarshalFunc = f
}

func (w *Client) SetAccessToken(token string) *Client {
	w.accessToken = token
	return w
}

func (w *Client) SetRefreshToken(token string) *Client {
	w.refreshToken = token
	return w
}

func (w *Client) NewRequest() *resty.Request {
	return w.client.R()
}
