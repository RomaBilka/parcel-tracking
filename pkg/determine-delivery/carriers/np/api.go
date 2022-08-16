package np

import (
	"encoding/json"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/valyala/fasthttp"
)

const URL = "/v2.0/json/"

var handleErrors = map[string]error{
	"20000500603": response_errors.NotFound,      //Not found
	"20000300415": response_errors.NotFound,      //Document not found
	"20000200161": response_errors.NotFound,      //Documents not found
	"20000200157": response_errors.InvalidNumber, //Document number empty
	"20000200158": response_errors.InvalidNumber, //Document number incorrect
	"20001401442": response_errors.InvalidNumber, //Document number is not correct
	"20001401158": response_errors.InvalidNumber, //Error: wrong format of the document number
	"20000202391": response_errors.InvalidNumber, //Invalid document number
	"20000401614": response_errors.InvalidNumber, //Wrong DocumentNumber
	"20000401620": response_errors.InvalidNumber, //DocumentNumber is invalid
	"20000201873": response_errors.InvalidNumber, //DocumentNumber invalid
	"20000201994": response_errors.InvalidNumber, //Invalid DocumentNumber
	"20002102188": response_errors.InvalidNumber, //DocumentNumber is incorrect
	"20001402465": response_errors.InvalidNumber, //DocumentNumber invalid format
	"20002102542": response_errors.InvalidNumber, //Invalid DocumentNumber format
	"20002102543": response_errors.InvalidNumber, //Document not found by DocumentNumber
	"20001402806": response_errors.InvalidNumber, //DocumentNumber invalid format.
	"20002502852": response_errors.InvalidNumber, //DocumentNumber is empty
}

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
		return nil, getErrors(trackingDocumentsResponse)
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

func getErrors(TrackingDocumentsResponse *TrackingDocumentsResponse) error {
	for _, code := range TrackingDocumentsResponse.ErrorCodes {
		if err, ok := handleErrors[code]; ok {
			return err
		}
	}

	lenErrors := len(TrackingDocumentsResponse.Errors)
	if lenErrors > 0 {
		errorMsgs := ""
		for i, e := range TrackingDocumentsResponse.Errors {
			errorMsgs += e
			if i < -1 {
				errorMsgs += ", "
			}
		}

		return errors.New(errorMsgs)
	}

	return nil
}
