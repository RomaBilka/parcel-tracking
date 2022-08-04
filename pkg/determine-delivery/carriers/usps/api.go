package usps

import (
	"errors"
	"fmt"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	"github.com/valyala/fasthttp"
)

type Api struct {
	UserId   string
	Password string
	URL      string
}

func NewApi(userId, password, url string) *Api {
	return &Api{
		UserId:   userId,
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
	if err = r.error(b); err != nil {
		return nil, err
	}

	if err = r.unmarshalTrackData(b); err != nil {
		return nil, err
	}
	return r, nil
}

func (api *Api) makeRequest(trackingNum, method, endPoint string) ([]byte, error) {
	xmlBody := fmt.Sprintf(`&XML=<TrackRequest USERID="%s" PASSWORD="%s"><TrackID ID="%s"></TrackID></TrackRequest>`,
		api.UserId, api.Password, trackingNum)

	res, err := http.Do(endPoint, method, func(req *fasthttp.Request) {
		req.Header.SetContentType(http.XmlContentType)
		req.AppendBodyString(xmlBody)
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
