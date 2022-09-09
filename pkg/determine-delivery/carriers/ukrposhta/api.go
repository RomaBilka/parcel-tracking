package ukrposhta

import (
	"encoding/json"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	"github.com/valyala/fasthttp"
)

const URL = "/statuses/"

type Api struct {
	apiURL string
	token  string
}

func NewApi(apiURL, token string) *Api {
	return &Api{
		apiURL: apiURL,
		token:  token,
	}
}

func (api *Api) TrackByTrackingNumber(barcodes []string) (*Response, error) {
	data, err := json.Marshal(barcodes)
	if err != nil {
		return nil, err
	}

	b, err := api.makeRequest(data, "with-not-found", fasthttp.MethodGet)
	if err != nil {
		return nil, err
	}

	response := &Response{}

	if err := json.Unmarshal(b, response); err != nil {
		return nil, err
	}

	return response, err
}

func (api *Api) makeRequest(data []byte, path, method string) ([]byte, error) {
	res, err := http.Do(api.apiURL+URL+path, method, func(req *fasthttp.Request) {
		req.SetBody(data)
		req.Header.SetContentType("application/json")
		req.Header.Add("authorization", "Bearer "+api.token)
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	if res.StatusCode() == fasthttp.StatusOK {
		return res.Body(), nil
	}
	return nil, errors.New(fasthttp.StatusMessage(res.StatusCode()))
}
