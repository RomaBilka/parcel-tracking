package dhl

import (
	"encoding/json"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

type Api struct {
	apiURL string
	apiKey string
}

func NewApi(apiURL, apiKey string) *Api {
	return &Api{
		apiURL: apiURL,
		apiKey: apiKey,
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

	url := api.apiURL + endPoint + "?" + v.Encode()
	res, err := http.Do(url, method, func(req *fasthttp.Request) {
		req.Header.Add("DHL-API-Key", api.apiKey)
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	body := res.Body()
	if res.StatusCode() == fasthttp.StatusOK {
		return body, nil
	}

	errorResponse := &errorResponse{}
	if err := json.Unmarshal(body, errorResponse); err != nil {
		return nil, err
	}
	return nil, errors.New(errorResponse.Detail)
}
