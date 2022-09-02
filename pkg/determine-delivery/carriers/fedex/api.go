package fedex

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/RomaBilka/parcel-tracking/pkg/http"
	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

const (
	tokenExpirationLimit = 1
	gzip                 = "gzip"
)

var handleErrors = map[string]error{
	"TRACKING.REFRENCEVALUE.INVALID":    response_errors.InvalidNumber,
	"TRACKING.TCNVALUE.EMPTY":           response_errors.InvalidNumber,
	"TRACKING.TRACKINGNUMBER.EMPTY":     response_errors.InvalidNumber,
	"TRACKING.TRACKINGNUMBER.INVALID":   response_errors.InvalidNumber,
	"TRACKING.TRACKINGNUMBER.NOTFOUND":  response_errors.NotFound,
	"NOTIFICATION.TRACKINGNBR.NOTFOUND": response_errors.NotFound,
	"TRACKING.REFERENCENUMBER.NOTFOUND": response_errors.NotFound,
	"TRACKING.TCN.NOTFOUND":             response_errors.NotFound,
}

type (
	Api struct {
		apiURL       string
		clientId     string
		grantType    string
		clientSecret string
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

func NewApi(apiURL, apiKey, grantType, shippingAccount string) *Api {
	return &Api{
		apiURL:       apiURL,
		clientId:     apiKey,
		grantType:    grantType,
		clientSecret: shippingAccount,
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
		contentType:       http.JsonContentType,
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
				if err := getHandleErrors(result.Error.Code); err != nil {
					return nil, err
				}
				return nil, errors.New(result.Error.Message)
			}
		}
	}

	return trackingResponse, err
}

func (api *Api) authorize() (string, error) {
	api.token.Lock()
	defer api.token.Unlock()

	if !isExpired(api.token.expire) {
		return api.token.token, nil
	}

	authParams := authorizeRequest{
		GrantType:    api.grantType,
		ClientId:     api.clientId,
		ClientSecret: api.clientSecret,
	}

	v, err := query.Values(authParams)
	if err != nil {
		return "", err
	}

	request := requestParam{
		body:        []byte(v.Encode()),
		path:        "/oauth/token",
		method:      fasthttp.MethodPost,
		contentType: http.FormContentType,
	}

	response, err := api.makeRequest(request)
	if err != nil {
		return "", err
	}

	authResp := &authResponse{}
	if err := json.Unmarshal(response, authResp); err != nil {
		return "", err
	}

	if authResp.AccessToken == "" {
		res := &Response{}
		if err := json.Unmarshal(response, res); err != nil {
			return "", err
		}
		if err = getErrors(res.Errors); err != nil {
			return "", err
		}
	}

	if isExpired(authResp.ExpiresIn.Time) {
		return "", errors.New("short expiration of the token")
	}

	api.token.token = authResp.AccessToken
	api.token.expire = authResp.ExpiresIn.Time

	return api.token.token, nil
}

func (api *Api) makeRequest(r requestParam) ([]byte, error) {
	var token string
	if r.needAuthorization {
		authToken, err := api.authorize()
		if err != nil {
			return nil, err
		}
		token = authToken
	}

	res, err := http.Do(api.apiURL+r.path, r.method, func(req *fasthttp.Request) {
		if r.needAuthorization {
			req.Header.Add("authorization", "Bearer "+token)
		}
		req.SetBody(r.body)
		req.Header.SetContentType(r.contentType)
	})
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseResponse(res)

	if string(res.Header.ContentType()) != gzip {
		return res.Body(), nil
	}

	b, err := res.BodyGunzip()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func getErrors(err []Error) error {
	lenErrors := len(err)
	if lenErrors > 0 {
		errorMsgs := ""
		for i, e := range err {
			if err := getHandleErrors(e.Code); err != nil {
				return err
			}

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

func getHandleErrors(code string) error {
	if err, ok := handleErrors[code]; ok {
		return err
	}
	return nil
}
