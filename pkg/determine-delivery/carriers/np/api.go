package np

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

const URL = "/v2.0/json/"

type novaPoshta struct {
	apiKey string
	apiURL string
}

func NewNovaPoshta(apiURL, apiKey string) *novaPoshta {
	return &novaPoshta{
		apiKey: apiKey,
		apiURL: apiURL,
	}
}

func (np *novaPoshta) TrackingDocument(methodProperties TrackingDocuments) ([]TrackingDocumentResponse, error) {
	req := novaPoshtaRequest{
		ModelName:    "TrackingDocument",
		CalledMethod: "getStatusDocuments",
	}
	req.MethodProperties = methodProperties

	trackingDocumentsResponse := &TrackingDocumentsResponse{}
	b, err := np.makeRequest(req, fasthttp.MethodGet)
	if err != nil {
		return []TrackingDocumentResponse{}, err
	}

	err = json.Unmarshal(b, trackingDocumentsResponse)
	if err != nil {
		return []TrackingDocumentResponse{}, err
	}

	return trackingDocumentsResponse.Data, err
}

func (np *novaPoshta) makeRequest(r novaPoshtaRequest, method string) ([]byte, error) {
	body := make([]byte, 0)
	r.ApiKey = np.apiKey

	data, err := json.Marshal(r)
	if err != nil {
		return body, err
	}

	req := fasthttp.AcquireRequest()
	req.SetBody(data)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")
	req.SetRequestURI(np.apiURL + URL)
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		return body, err
	}
	fasthttp.ReleaseRequest(req)
	body = res.Body()

	return body, nil
}