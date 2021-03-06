package sdreq

import (
	"github.com/gaorx/stardust4/sderr"
	"github.com/imroc/req/v3"
)

func PostForResponse(client *req.Client, url string, body any, opts ...RequestOption) (*req.Response, error) {
	if client == nil {
		client = req.DefaultClient()
	}
	request := applyOptions(client.R(), opts).SetBody(body)
	response, err := request.Post(url)
	if err != nil {
		return nil, sderr.Wrap(err, "sdreq get response error")
	}
	return response, nil
}

func PostForBytes(client *req.Client, url string, body any, opts ...RequestOption) (int, []byte, error) {
	response, err := PostForResponse(client, url, body, opts...)
	if err != nil {
		return 0, nil, err
	}
	data, err := response.ToBytes()
	if err != nil {
		return response.StatusCode, nil, sderr.Wrap(err, "sdreq response to bytes error")
	}
	return response.StatusCode, data, nil
}

func PostForText(client *req.Client, url string, body any, opts ...RequestOption) (int, string, error) {
	response, err := PostForResponse(client, url, body, opts...)
	if err != nil {
		return 0, "", err
	}
	data, err := response.ToString()
	if err != nil {
		return response.StatusCode, "", sderr.Wrap(err, "sdreq response to text error")
	}
	return response.StatusCode, data, nil
}

func PostForJson[R any](client *req.Client, url string, body any, opts ...RequestOption) (int, R, error) {
	var r R
	response, err := PostForResponse(client, url, body, opts...)
	if err != nil {
		return 0, r, err
	}
	err = response.UnmarshalJson(&r)
	if err != nil {
		return response.StatusCode, r, sderr.Wrap(err, "sdreq response unmarshal json error")
	}
	return response.StatusCode, r, nil
}
