package ups

import (
	"encoding/xml"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/valyala/fasthttp"
)

var handleErrors = map[string]error{
	"150022": response_errors.InvalidNumber, //Invalid tracking number
	"150028": response_errors.InvalidNumber, //Hard Invalid shipper number length
	"151018": response_errors.InvalidNumber, //Invalid tracking number
	"151044": response_errors.NotFound,      //No tracking information available
	"151045": response_errors.NotFound,      //No information found
	"151062": response_errors.NotFound,      //No tracking information available
	"151068": response_errors.InvalidNumber, //Invalid Shipper Number
	"152110": response_errors.NotFound,      //No information found for reference number
	"154030": response_errors.NotFound,      //No information for this tracking number

}

type Api struct {
	apiURL              string
	userId              string
	accessLicenseNumber string
	password            string
}

func NewApi(apiURL, userId, accessLicenseNumber, password string) *Api {
	return &Api{
		apiURL:              apiURL,
		userId:              userId,
		accessLicenseNumber: accessLicenseNumber,
		password:            password,
	}
}

func (api *Api) TrackByTrackingNumber(trackingNumber string) (*TrackResponse, error) {
	trackRequest := TrackRequest{TrackingNumber: trackingNumber}
	b, err := api.makeRequest(trackRequest, "/Track")
	if err != nil {
		return nil, err
	}

	trackResponse := &TrackResponse{}
	if err := xml.Unmarshal(b, trackResponse); err != nil {
		return nil, err
	}

	if trackResponse.Response.Error.ErrorCode != "" {
		if err, ok := handleErrors[trackResponse.Response.Error.ErrorCode]; ok {
			return nil, err
		}

		return nil, errors.New(trackResponse.Response.Error.ErrorDescription)
	}

	return trackResponse, nil
}

func (api *Api) makeRequest(r TrackRequest, path string) ([]byte, error) {
	accessRequest := AccessRequest{
		AccessLicenseNumber: api.accessLicenseNumber,
		UserId:              api.userId,
		Password:            api.password,
	}

	accessByte, err := xml.MarshalIndent(accessRequest, "", " ")
	if err != nil {
		return nil, err
	}
	data := append([]byte(xml.Header), accessByte...)

	requestByte, err := xml.MarshalIndent(r, "", " ")
	if err != nil {
		return nil, err
	}
	data = append(data, []byte(xml.Header)...)
	data = append(data, requestByte...)

	res, err := http.Do(api.apiURL+path, fasthttp.MethodPost, func(req *fasthttp.Request) {
		req.SetBody(data)
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	return res.Body(), nil
}
