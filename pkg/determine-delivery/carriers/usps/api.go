package usps

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

type Api struct {
	UserID   string
	Password string
	URL      string
}

func NewApi(userID, password, url string) *Api {
	return &Api{
		UserID:   userID,
		Password: password,
		URL:      url,
	}
}

func (api *Api) TrackingDocument(trackNumber string) (*response, error) {
	b, err := api.makeRequest(trackNumber, fasthttp.MethodPost, api.URL)

	if err != nil {
		return nil, err
	}

	r := &response{}

	if err = r.isError(b); err != nil {
		return nil, err
	}

	if err = r.unmarshalTrackData(b); err != nil {
		return nil, err
	}

	return r, nil
}

const xmlBodyRequest = "&XML=<TrackRequest USERID=\"%s\" PASSWORD=\"%s\"><TrackID ID=\"%s\"></TrackID></TrackRequest>"

// USPS for tracking request use the POST request
func (api *Api) makeRequest(trackingNum, method, endPoint string) ([]byte, error) {

	xmlBody := fmt.Sprintf(xmlBodyRequest, api.UserID, api.Password, trackingNum)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("text/xml")
	req.Header.SetMethod(method)
	req.SetRequestURI(endPoint)
	req.AppendBodyString(xmlBody)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.Do(req, res); err != nil {
		log.Fatalf("Response failed with status code: %d and body: %s\n", res.StatusCode(), res.Body())
	}

	if res.StatusCode() == fasthttp.StatusOK {
		return res.Body(), nil
	}

	return nil, errors.New(fasthttp.StatusMessage(res.StatusCode()))
}
