package sdk

import "resty.dev/v3"

type RestyOption func(request *resty.Request)

func ReqWithJson(json Json) RestyOption {
	return func(request *resty.Request) {
		request.
			SetHeader("Content-Type", "application/json").
			SetBody(json)
	}
}

func ReqWithForm(form Form) RestyOption {
	return func(request *resty.Request) {
		request.
			SetFormData(form)
	}
}

func ReqWithResp(v any) RestyOption {
	return func(request *resty.Request) {
		request.SetResult(v)
	}
}

func ReqWithQuery(query Form) RestyOption {
	return func(request *resty.Request) {
		request.SetQueryParams(query)
	}
}
