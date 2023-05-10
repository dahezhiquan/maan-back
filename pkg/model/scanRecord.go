package model

import "gorm.io/gorm"

// 表名称

const TableNameScanRecord = "scan_record"

type ScanRecord struct {
	gorm.Model
	Content      string `gorm:"column:content;type:longtext;not null;comment:解析内容" json:"content"`        // 解析内容
	Mvss         int32  `gorm:"column:mvss;type:int;comment:MVSS分数" json:"mvss"`                          // MVSS分数
	IpAddr       string `gorm:"column:ip_addr;type:varchar(100);comment:IP归属地" json:"ip_addr"`            // 解析ip
	RiskType     string `gorm:"column:risk_type;type:varchar(100);comment:风险类型" json:"risk_type"`         // 风险类型
	ContentTitle string `gorm:"column:content_title;type:varchar(200);comment:扫描标题" json:"content_title"` // 扫描标题
	IsSafe       int32  `gorm:"column:is_safe;type:int;comment:是否安全" json:"is_safe"`                      // 是否安全（1 - 安全 2 - 警告 3 - 危险）
	IsDfa        int32  `gorm:"column:is_dfa;type:int;comment:是否涉黄涉黑涉政" json:"is_dfa"`                    // 是否涉黄涉黑涉政（1 - 涉及 0 - 不涉及）
}

func (*ScanRecord) TableName() string {
	return TableNameScanRecord
}
