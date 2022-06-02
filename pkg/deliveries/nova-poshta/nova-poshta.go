package nova_poshta

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type responseInterface interface {
	getError() error
}

type novaPoshta struct {
	apiKey     string
	requestURL string
}

func NewNovaPoshta(requestURL, apiKey string) *novaPoshta {
	return &novaPoshta{
		apiKey:     apiKey,
		requestURL: requestURL,
	}
}

func (np *novaPoshta) TrackingDocument(methodProperties TrackingDocuments) ([]TrackingDocumentResponse, error) {
	req := novaPoshtaRequest{
		ModelName:    "TrackingDocument",
		CalledMethod: "getStatusDocuments",
	}
	req.MethodProperties = methodProperties

	res := &TrackingDocumentsResponse{}
	_, err := np.makeRequest(req, res, http.MethodGet)
	if err != nil {
		return []TrackingDocumentResponse{}, err
	}

	err = res.getError()
	if err != nil {
		return []TrackingDocumentResponse{}, err
	}

	return res.Data, err
}

func (np *novaPoshta) makeRequest(r novaPoshtaRequest, npResponse responseInterface, method string) (responseInterface, error) {
	r.ApiKey = np.apiKey

	data, err := json.Marshal(r)
	if err != nil {
		return npResponse, err
	}
	reader := bytes.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, np.requestURL, reader)

	if err != nil {
		return npResponse, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return npResponse, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return npResponse, err
	}

	err = json.Unmarshal(b, npResponse)
	if err != nil {
		return npResponse, err
	}

	return npResponse, nil
}
