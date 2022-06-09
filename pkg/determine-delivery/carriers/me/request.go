package me

type param struct {
	meestExpressRequest `xml:"param"`
}

type meestExpressRequest struct {
	Login    string `xml:"login"`
	Function string `xml:"function"`
	Where    string `xml:"where"`
	Order    string `xml:"order"`
	Sign     string `xml:"sign"`
}
