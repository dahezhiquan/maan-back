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
	IsPassDfa int    `json:"is_pass_dfa"`
	UrlTitle  string `json:"url_title"`
}

type IntelligenceResp struct {
	IntelligenceList []*Intelligence `json:"intelligence_list"`
}

type Intelligence struct {
	ID       int64  `json:"id"`
	Content  string `json:"content"`   // 解析内容
	Mvss     int32  `json:"mvss"`      // MVSS分数
	RiskType string `json:"risk_type"` // 风险类型
	IsSafe   int32  `json:"is_safe"`   // 是否安全（1 - 安全 2 - 警告 3 - 危险）
	IsDfa    int32  `json:"is_dfa"`    // 是否涉黄涉黑涉政（1 - 涉及 0 - 不涉及）
}
