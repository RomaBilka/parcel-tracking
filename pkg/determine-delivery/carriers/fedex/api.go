package fedex

type Api struct {
	grantType    string
	clientId     string
	clientSecret string
	apiURL       string
}

func NewApi(grantType, clientId, clientSecret, apiURL string) *Api {
	return &Api{
		grantType:    grantType,
		clientId:     clientId,
		clientSecret: clientSecret,
		apiURL:       apiURL,
	}
}

func (api *Api) TrackByTrackingNumber(trackNumber string) {

}

/*
func (api *Api) ShipmentsTrack(trackNumber string) (*ShipmentsTrackResponse, error) {
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

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetBody(data)
	req.Header.SetMethod(method)
	req.Header.SetContentType("text/xml")
	req.SetRequestURI(api.apiURL)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	return res.Body(), nil
}

func (me *Api) getHash(r meestExpressRequest) string {
	hash := md5.Sum([]byte(r.Login + me.password + r.Function + r.Where + r.Order))

	return hex.EncodeToString(hash[:])
}
*/
