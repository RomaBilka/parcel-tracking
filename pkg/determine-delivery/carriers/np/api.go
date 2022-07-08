package np

import (
	"encoding/json"
	"errors"

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

func (np *Api) TrackingDocument(methodProperties TrackingDocuments) (*TrackingDocumentsResponse, error) {
	req := novaPoshtaRequest{
		ModelName:    "TrackingDocument",
		CalledMethod: "getStatusDocuments",
	}
	req.MethodProperties = methodProperties

	trackingDocumentsResponse := &TrackingDocumentsResponse{}
	b, err := np.makeRequest(req, fasthttp.MethodGet)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, trackingDocumentsResponse)
	if err != nil {
		return nil, err
	}

	if !trackingDocumentsResponse.Success {
		return nil, errors.New(trackingDocumentsResponse.Errors[0])
	}

	return trackingDocumentsResponse, err
}

func (np *Api) makeRequest(r novaPoshtaRequest, method string) ([]byte, error) {
	r.ApiKey = np.apiKey

	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetBody(data)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")
	req.SetRequestURI(np.apiURL + URL)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	return res.Body(), nil
}
