package http

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	XmlContentType  = "text/xml"
	JsonContentType = "application/json"
	FormContentType = "application/x-www-form-urlencoded"
)

func Do(url, method string, upd func(*fasthttp.Request)) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	if upd != nil {
		upd(req)
	}

	res := fasthttp.AcquireResponse()
	if err := fasthttp.DoTimeout(req, res, 60*time.Second); err != nil {
		defer fasthttp.ReleaseResponse(res)
		return nil, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode(), res.Body())
	}
	return res, nil
}
