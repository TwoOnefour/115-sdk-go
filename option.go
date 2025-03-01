package _115

import "resty.dev/v3"

type Option func(*Client)

func WithJsonMarshalFunc(f func(v interface{}) ([]byte, error)) Option {
	return func(w *Client) {
		w.SetJsonMarshalFunc(f)
	}
}

func WithJsonUnmarshalFunc(f func(data []byte, v interface{}) error) Option {
	return func(w *Client) {
		w.SetJsonUnmarshalFunc(f)
	}
}

func WithRestyClient(rc *resty.Client) Option {
	return func(c *Client) {
		c.client = rc
	}
}

func WithDebug() Option {
	return func(c *Client) {
		c.SetDebug(true)
	}
}

func WithTrace() Option {
	return func(c *Client) {
		c.EnableTrace()
	}
}

func WithProxy(proxy string) Option {
	return func(c *Client) {
		c.SetProxy(proxy)
	}
}

func WithAccessToken(token string) Option {
	return func(w *Client) {
		w.SetAccessToken(token)
	}
}

func WithRefreshToken(token string) Option {
	return func(w *Client) {
		w.SetRefreshToken(token)
	}
}

type RestyOption func(request *resty.Request)
