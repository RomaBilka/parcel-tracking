package usps

import (
	"encoding/xml"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	"github.com/valyala/fasthttp"
)

type Api struct {
	apiURL   string
	userId   string
	sourceId string
}

func NewApi(apiURL, userId, sourceId string) *Api {
	return &Api{
		apiURL:   apiURL,
		userId:   userId,
		sourceId: sourceId,
	}
}

func (api *Api) TrackByTrackingNumber(trackNumbers []TrackID) (*TrackResponse, error) {
	t := TrackFieldRequest{
		Revision: 1,
		TrackID:  trackNumbers,
	}

	b, err := api.makeRequest(t, fasthttp.MethodPost, api.apiURL)
	if err != nil {
		return nil, err
	}

	trackResponse := &TrackResponse{}
	if err := xml.Unmarshal(b, trackResponse); err != nil {
		return nil, err
	}

	return trackResponse, nil
}

func (api *Api) makeRequest(t TrackFieldRequest, method, endPoint string) ([]byte, error) {
	requestByte, err := xml.MarshalIndent(t, "", " ")
	if err != nil {
		return nil, err
	}

	data := append([]byte(xml.Header), requestByte...)

	res, err := http.Do(endPoint, method, func(req *fasthttp.Request) {
		req.Header.SetContentType(http.XmlContentType)
		req.SetBody(data)
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
