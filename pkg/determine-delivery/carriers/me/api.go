package me

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"errors"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/valyala/fasthttp"
)

const statusOk = "000" //Ok

var handleErrors = map[string]error{
	"103": response_errors.NotFound, //Document not found
	"104": response_errors.NotFound, //Directory not found
}

type Api struct {
	apiURL   string
	agentUID string
	login    string
	password string
}

func NewApi(apiURL, agentUID, login, password string) *Api {
	return &Api{
		apiURL:   apiURL,
		agentUID: agentUID,
		login:    login,
		password: password,
	}
}

func (api *Api) TrackByTrackingNumber(trackNumber string) (*ShipmentsTrackResponse, error) {
	req := meestExpressRequest{
		Function: "SHIPMENTS_TRACK",
		Where:    api.agentUID + "," + trackNumber,
	}

	b, err := api.makeRequest(req, fasthttp.MethodPost)
	if err != nil {
		return nil, err
	}
	shipmentsTrackResponse := &ShipmentsTrackResponse{}

	if err := xml.Unmarshal(b, shipmentsTrackResponse); err != nil {
		return nil, err
	}

	if shipmentsTrackResponse.Errors.Code != statusOk {

		if err, ok := handleErrors[shipmentsTrackResponse.Errors.Code]; ok {
			return nil, err
		}
		return nil, errors.New(shipmentsTrackResponse.Errors.Name)
	}

	return shipmentsTrackResponse, nil
}

func (api *Api) makeRequest(r meestExpressRequest, method string) ([]byte, error) {
	r.Login = api.login
	r.Sign = api.getHash(r)
	p := param{r}

	xmlString, _ := xml.MarshalIndent(p, "", " ")
	data := append([]byte(xml.Header), xmlString...)

	res, err := http.Do(api.apiURL, method, func(req *fasthttp.Request) {
		req.SetBody(data)
		req.Header.SetContentType(http.XmlContentType)
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	return res.Body(), nil
}

func (me *Api) getHash(r meestExpressRequest) string {
	hash := md5.Sum([]byte(r.Login + me.password + r.Function + r.Where + r.Order))

	return hex.EncodeToString(hash[:])
}
