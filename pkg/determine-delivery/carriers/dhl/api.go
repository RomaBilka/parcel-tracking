package dhl

import (
	"encoding/json"
	"errors"

	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

type Api struct {
	apiKey string
	apiURL string
}

func NewApi(apiURL, apiKey string) *Api {
	return &Api{
		apiKey: apiKey,
		apiURL: apiURL,
	}
}

func (api *Api) TrackingDocument(trackNumber string) (*response, error) {
	req := request{
		TrackingNumber: trackNumber,
	}

	b, err := api.makeRequest(req, fasthttp.MethodGet, "/track/shipments")

	if err != nil {
		return nil, err
	}

	r := &response{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}

	return r, nil
}

func (api *Api) makeRequest(r request, method, endPoint string) ([]byte, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)
	req.Header.Add("DHL-API-Key", api.apiKey)
	req.SetRequestURI(api.apiURL + endPoint + "?" + v.Encode())
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	fasthttp.ReleaseRequest(req)
	body := res.Body()

	if res.StatusCode() == fasthttp.StatusOK {
		return body, nil
	}

	return nil, errors.New(fasthttp.StatusMessage(res.StatusCode()))
}
