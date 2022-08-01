package fedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

const (
	tokenExpirationLimit = 1
	gzip                 = "gzip"
)

type (
	Api struct {
		apiURL       string
		grantType    string //client_credentials
		clientId     string //API KEY
		clientSecret string //SHIPPING ACCOUNT
		token        struct {
			token  string
			expire time.Time
			sync.Mutex
		}
	}

	requestParam struct {
		body              []byte
		path              string
		method            string
		contentType       string
		needAuthorization bool
	}
)

func NewApi(apiURL, grantType, clientId, clientSecret string) *Api {
	return &Api{
		apiURL:       apiURL,
		grantType:    grantType,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (api *Api) TrackByTrackingNumber(trackingRequest TrackingRequest) (*TrackingResponse, error) {
	b, err := json.Marshal(trackingRequest)
	if err != nil {
		return nil, err
	}

	request := requestParam{
		body:              b,
		path:              "/track/v1/trackingdocuments",
		method:            fasthttp.MethodPost,
		contentType:       "application/json",
		needAuthorization: true,
	}

	response, err := api.makeRequest(request)
	if err != nil {
		return nil, err
	}

	trackingResponse := &TrackingResponse{}
	if err := json.Unmarshal(response, trackingResponse); err != nil {
		return nil, err
	}

	if err = getErrors(trackingResponse.Errors); err != nil {
		return nil, err
	}

	for _, rules := range trackingResponse.Output.CompleteTrackResults {
		for _, result := range rules.TrackResults {
			if result.Error.Code != "" {
				return nil, errors.New(result.Error.Message)
			}
		}
	}

	return trackingResponse, err
}

func (api *Api) authorize() error {
	api.token.Lock()
	defer api.token.Unlock()

	authParams := authorizeRequest{
		GrantType:    api.grantType,
		ClientId:     api.clientId,
		ClientSecret: api.clientSecret,
	}

	v, err := query.Values(authParams)
	if err != nil {
		return err
	}

	request := requestParam{
		body:        []byte(v.Encode()),
		path:        "/oauth/token",
		method:      fasthttp.MethodPost,
		contentType: "application/x-www-form-urlencoded",
	}

	response, err := api.makeRequest(request)
	if err != nil {
		return err
	}

	authResp := &authResponse{}
	if err := json.Unmarshal(response, authResp); err != nil {
		fmt.Println(err)
		return err
	}

	if authResp.AccessToken == "" {
		res := &Response{}
		if err := json.Unmarshal(response, res); err != nil {
			return err
		}
		if err = getErrors(res.Errors); err != nil {
			return err
		}
	}

	if isExpired(authResp.ExpiresIn.Time) {
		return errors.New("short expiration of the token")
	}

	api.token.token = authResp.AccessToken
	api.token.expire = authResp.ExpiresIn.Time

	return nil
}

func (api *Api) makeRequest(r requestParam) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetBody(r.body)
	req.Header.SetMethod(r.method)
	req.Header.SetContentType(r.contentType)
	req.SetRequestURI(api.apiURL + r.path)

	if r.needAuthorization {
		if err := api.setAuthorize(req); err != nil {
			return nil, err
		}
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	if string(res.Header.ContentEncoding()) == gzip {
		b, err := res.BodyGunzip()
		if err != nil {
			return nil, err
		}

		return b, nil
	}

	return res.Body(), nil
}

func (api *Api) setAuthorize(req *fasthttp.Request) error {
	if isExpired(api.token.expire) {
		if err := api.authorize(); err != nil {
			return err
		}
	}
	req.Header.Add("authorization", "Bearer "+api.token.token)
	return nil
}

func getErrors(err []Error) error {
	lenErrors := len(err)
	if lenErrors > 0 {
		errorMsgs := ""
		for i, e := range err {
			errorMsgs += e.Message
			if i < lenErrors-1 {
				errorMsgs += ", "
			}
		}
		return errors.New(errorMsgs)
	}
	return nil
}

func isExpired(t time.Time) bool {
	return time.Until(t) < tokenExpirationLimit*time.Minute
}
