package fuckqr

import (
	"maan/pkg/dao"
)

// 直接从百万计ip库中查找

func HackerIpCheck(ip string) bool {
	d := dao.NewIpSumDao()
	ipSum, _ := d.FindIpSumByIp(nil, nil, ip)
	if ipSum.HackerIp == "" {
		return true
	} else {
		return false
	}
}
