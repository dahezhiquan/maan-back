package fuckqr

import (
	"maan/pkg/dto"
	"maan/pkg/public"
	"strings"
)

func RiskAnalyzer(resp dto.QrScanResp) string {
	if resp.Mvss >= 60 {
		return ""
	}
	if strings.Contains(resp.UrlTitle, "闲鱼") {
		return public.XianYuRisk
	}
	if strings.Contains(resp.UrlTitle, "支付") {
		return public.PayRisk
	}
	return ""
}
