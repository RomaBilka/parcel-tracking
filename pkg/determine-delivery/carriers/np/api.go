package np

import (
	"encoding/json"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	"github.com/valyala/fasthttp"
)

const URL = "/v2.0/json/"

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

func (api *Api) TrackByTrackingNumber(methodProperties TrackingDocuments) (*TrackingDocumentsResponse, error) {
	req := novaPoshtaRequest{
		ModelName:    "TrackingDocument",
		CalledMethod: "getStatusDocuments",
	}
	req.MethodProperties = methodProperties

	trackingDocumentsResponse := &TrackingDocumentsResponse{}
	b, err := api.makeRequest(req, fasthttp.MethodGet)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, trackingDocumentsResponse); err != nil {
		return nil, err
	}

	if !trackingDocumentsResponse.Success {
		return nil, errors.New(trackingDocumentsResponse.Errors[0])
	}

	return trackingDocumentsResponse, err
}

func (api *Api) makeRequest(r novaPoshtaRequest, method string) ([]byte, error) {
	r.ApiKey = api.apiKey
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	res, err := http.Do(api.apiURL+URL, method, func(req *fasthttp.Request) {
		req.SetBody(data)
		req.Header.SetContentType("application/json")
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	return res.Body(), nil
}
