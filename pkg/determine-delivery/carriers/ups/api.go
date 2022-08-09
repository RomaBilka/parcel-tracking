package ups

import (
	"encoding/xml"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	"github.com/valyala/fasthttp"
)

type Api struct {
	apiURL              string
	accessLicenseNumber string
	userId              string
	password            string
}

func NewApi(apiURL, accessLicenseNumber, userId, password string) *Api {
	return &Api{
		apiURL:              apiURL,
		accessLicenseNumber: accessLicenseNumber,
		userId:              userId,
		password:            password,
	}
}

func (api *Api) TrackByNumber(trackingNumber string) (*TrackResponse, error) {
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

	xmlString, err := xml.MarshalIndent(accessRequest, "", " ")
	if err != nil {
		return nil, err
	}
	data := append([]byte(xml.Header), xmlString...)

	xmlString, err = xml.MarshalIndent(r, "", " ")
	if err != nil {
		return nil, err
	}
	data = append(data, []byte(xml.Header)...)
	data = append(data, xmlString...)

	res, err := http.Do(api.apiURL+path, fasthttp.MethodPost, func(req *fasthttp.Request) {
		req.SetBody(data)
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	return res.Body(), nil
}
