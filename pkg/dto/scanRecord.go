package dto

// 前端传递的二维码解析数据 req

type QrScanReq struct {
	Content string `json:"content" form:"content" validate:"max=50000"`
}

// 解析完成的二维码数据 resp

type QrScanResp struct {
	Mvss      int    `json:"mvss"`
	IpAddr    string `json:"ip_addr"`
	RiskType  string `json:"risk_type"`
	IsPassDfa bool   `json:"is_pass_dfa"`
	UrlTitle  string `json:"url_title"`
}
