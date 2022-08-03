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

	upd(req)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.DoTimeout(req, res, time.Second*60); err != nil {
		return nil, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode(), res.Body())
	}

	return res, nil
}
