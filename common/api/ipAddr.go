package api

import (
	"encoding/json"
	"io"
	"maan/common/errs"
	"maan/pkg/public"
	"net/http"
)

type IpAddrInfo struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Ipinfo Ipinfo `json:"ipinfo"`
	Ipdata Ipdata `json:"ipdata"`
	Adcode Adcode `json:"adcode"`
	Tips   string `json:"tips"`
	Time   int    `json:"time"`
}
type Ipinfo struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Cnip bool   `json:"cnip"`
}
type Ipdata struct {
	Info1 string `json:"info1"`
	Info2 string `json:"info2"`
	Info3 string `json:"info3"`
	Isp   string `json:"isp"`
}
type Adcode struct {
	O string      `json:"o"`
	P string      `json:"p"`
	C string      `json:"c"`
	N string      `json:"n"`
	R interface{} `json:"r"`
	A string      `json:"a"`
	I bool        `json:"i"`
}

// 获取ip归属地信息

func GetIpInfo(ip string) (ipAddrInfo IpAddrInfo, error *errs.BError) {
	resp, err := http.Get(public.IPADDRAPI + ip)
	if err != nil {
		return ipAddrInfo, errs.IpAddrSearchError
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&ipAddrInfo)
	if err != nil {
		return ipAddrInfo, errs.StrToJsonError
	}
	return ipAddrInfo, nil
}
