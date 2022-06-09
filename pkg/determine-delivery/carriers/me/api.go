package me

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/valyala/fasthttp"
)

type meestExpress struct {
	agentUID string
	login    string
	password string
	apiURL   string
}

func NewMeestExpress(agentUID, login, password, apiURL string) *meestExpress {
	return &meestExpress{
		agentUID: agentUID,
		login:    login,
		password: password,
		apiURL:   apiURL,
	}
}

func (me *meestExpress) ShipmentsTrack(trackNumber string) {
	req := meestExpressRequest{
		Function: "SHIPMENTS_TRACK",
		Where:    me.agentUID + "," + trackNumber,
	}

	b, err := me.makeRequest(req, fasthttp.MethodPost)
	//shipmentsTrackResponse:=ShipmentsTrackResponse{}
	//xml.Unmarshal(b, shipmentsTrackResponse)
	//fmt.Println(shipmentsTrackResponse)
	fmt.Println(b, err)
	_ = ioutil.WriteFile("test.xml", b, 0644)
}

func (me *meestExpress) makeRequest(r meestExpressRequest, method string) ([]byte, error) {
	body := make([]byte, 0)

	r.Login = me.login
	r.Sign = me.getHash(r)
	p := param{r}

	xmlString, _ := xml.MarshalIndent(p, "", " ")
	data:= append([]byte(xml.Header) , xmlString...)

	req := fasthttp.AcquireRequest()
	req.SetBody(data)
	req.Header.SetMethod(method)
	req.Header.SetContentType("text/xml")
	req.SetRequestURI(me.apiURL)
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		return body, err
	}

	fasthttp.ReleaseRequest(req)
	body = res.Body()

	return body, nil
}
func (me *meestExpress) getHash(r meestExpressRequest) string {
	hash := md5.Sum([]byte(r.Login + me.password + r.Function + r.Where + r.Order))

	return hex.EncodeToString(hash[:])
}
