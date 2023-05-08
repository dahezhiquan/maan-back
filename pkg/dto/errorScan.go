package dto

// 前端传来的误报数据

type ErrorScanReq struct {
	IpAddr       string `json:"ip_addr" validate:"max=100"`       // 解析ip
	RiskType     string `json:"risk_type" validate:"max=100"`     // 风险类型
	ContentTitle string `json:"content_title" validate:"max=100"` // 解析url标题
	IsDfa        int32  `json:"is_dfa"`                           // 是否涉黄涉政涉黑
	Content      string `json:"content" validate:"max=10000"`     // 解析内容
	Mvss         int32  `json:"mvss"`                             // mvss分数
}
