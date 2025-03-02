package sdk

import (
	"context"
	"encoding/json"
	"fmt"

	"resty.dev/v3"
)

func (c *Client) Request(ctx context.Context, url string, method string, opts ...RestyOption) (*resty.Response, error) {
	req := c.NewRequest(ctx)
	for _, opt := range opts {
		opt(req)
	}
	return req.Execute(method, url)
}

func (c *Client) passportRequest(ctx context.Context, url, method string, respData any, opts ...RestyOption) (*resty.Response, error) {
	var resp AuthResp[any]
	if respData != nil {
		resp.Data = respData
	}
	response, err := c.Request(ctx, url, method, append(opts, ReqWithResp(&resp))...)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return response, fmt.Errorf("code: %d, message: %s", resp.Code, resp.Message)
	}
	if resp.Error != "" {
		return response, fmt.Errorf("error: %s, errno: %d", resp.Error, resp.Errno)
	}
	return response, nil
}

func (c *Client) authRequest(ctx context.Context, url, method string, respData any, retry bool, opts ...RestyOption) (*resty.Response, error) {
	var resp Resp[json.RawMessage]
	response, err := c.Request(ctx, url, method, append(opts, ReqWithResp(&resp), func(request *resty.Request) {
		request.SetAuthToken(c.accessToken)
	})...)
	if err != nil {
		return nil, err
	}
	if !resp.State {
		if !retry && resp.Code == 99 {
			_, err := c.RefreshToken(ctx)
			if err != nil {
				return response, err
			}
			return c.authRequest(ctx, url, method, respData, true, opts...)
		}
		return response, &Error{Code: resp.Code, Message: resp.Message}
	}
	if respData != nil {
		err = json.Unmarshal(resp.Data, respData)
		if err != nil {
			return response, err
		}
	}
	return response, nil
}

func (c *Client) AuthRequest(ctx context.Context, url, method string, respData any, opts ...RestyOption) (*resty.Response, error) {
	return c.authRequest(ctx, url, method, respData, false, opts...)
}
