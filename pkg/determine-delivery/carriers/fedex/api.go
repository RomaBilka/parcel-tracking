package fedex

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

const tokenExpirationLimit = 1 //1 second
const gzip = "gzip"

type Api struct {
	grantType    string
	clientId     string
	clientSecret string
	apiURL       string
}

type requestParam struct {
	body              []byte
	path              string
	method            string
	contentType       string
	needAuthorization bool
}

type token struct {
	token  string
	expire time.Time
}

var t token

func NewApi(grantType, clientId, clientSecret, apiURL string) *Api {
	return &Api{
		grantType:    grantType,
		clientId:     clientId,
		clientSecret: clientSecret,
		apiURL:       apiURL,
	}
}

func (api *Api) TrackByTrackingNumber(trackingRequest TrackingRequest) (*Response, error) {
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

	res := &Response{}
	if err := json.Unmarshal(response, res); err != nil {
		return nil, err
	}

	if err = getErrors(res.Errors); err != nil {
		return nil, err
	}

	return res, err
}

func (api *Api) authorize() error {
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

	b, err := api.makeRequest(request)
	if err != nil {
		return err
	}

	a := &authResponse{}
	if err := json.Unmarshal(b, a); err != nil {

		return err
	}

	if a.AccessToken == "" {
		res := &Response{}
		if err := json.Unmarshal(b, res); err != nil {
			return err
		}
		if err = getErrors(res.Errors); err != nil {
			return err
		}
	}

	if a.ExpiresIn <= tokenExpirationLimit {
		return errors.New("short expiration of the token")
	}

	seconds := time.Duration(a.ExpiresIn) * time.Second
	t = token{
		token:  a.AccessToken,
		expire: time.Now().Local().Add(seconds),
	}

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
	if t.isExpired() {
		if err := api.authorize(); err != nil {
			return err
		}
	}
	req.Header.Add("authorization", "Bearer 1"+t.token)
	return nil
}

func getErrors(err Errors) error {
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

func (t *token) isExpired() bool {
	return t.expire.Sub(time.Now()) < 1*time.Minute
}
