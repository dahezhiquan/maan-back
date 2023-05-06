package fuckqr

import (
	"strings"
)

// 如果二维码解析的地址是国外的ip，那么很大概率为风险二维码

func IpAddrCheck(ipAddr string) bool {
	safeAddr := []string{"河北", "山西", "辽宁", "吉林", "黑龙江", "江苏", "浙江", "安徽", "福建", "江西", "山东", "河南",
		"湖北", "湖南", "广东", "海南", "四川", "贵州", "云南", "陕西", "甘肃", "青海",
		"北京", "天津", "上海", "重庆", "内蒙古", "广西", "西藏", "宁夏", "新疆"}

	var flag bool = false
	for _, s := range safeAddr {
		if strings.Contains(ipAddr, s) {
			flag = true
			break
		}
	}

	// 国外服务加白
	safeServer := []string{"GitHub", "腾讯"}
	for _, s := range safeServer {
		if strings.Contains(ipAddr, s) {
			flag = true
			break
		}
	}

	return flag
}
