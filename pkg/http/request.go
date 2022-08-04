package http

import (
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
		return nil, err
	}
	return res, nil
}
