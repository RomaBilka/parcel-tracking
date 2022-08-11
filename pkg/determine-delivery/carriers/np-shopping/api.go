package np_shopping

import (
	"encoding/json"
	"fmt"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	"github.com/valyala/fasthttp"
)

type Api struct {
	url string
}

func NewApi() *Api {
	return &Api{url: "https://novaposhtaglobal.ua/ajax.php"}
}

func (api *Api) TrackByTrackingNumber(trackingID string) (*TrackingDocumentResponse, error) {
	// ajax_tracking is copy paste from https://novaposhtaglobal.ua/track/?Tracking_ID=3234 network call.
	// I can't find any documentation about it, so leaving it as in the request.
	// If you can find a documentation, or just you know other way to implement it, please let me know.
	action := "ajax_tracking"
	url := fmt.Sprintf("%s?Tracking_ID=%s&action=%s", api.url, trackingID, action)
	res, err := http.Do(url, fasthttp.MethodGet, func(req *fasthttp.Request) {
		req.Header.SetContentType("application/json")
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	body := res.Body()
	trackingDocumentsResponse := &TrackingDocumentResponse{}
	if err := json.Unmarshal(body, trackingDocumentsResponse); err != nil {
		return nil, err
	}

	if trackingDocumentsResponse.State == "" && trackingDocumentsResponse.Result != "" {
		return nil, fmt.Errorf("document number is not correct: %s", body)
	}
	if trackingDocumentsResponse.Result == "" {
		return nil, fmt.Errorf("something went wrong: %s", body)
	}

	return trackingDocumentsResponse, err
}
