package fuckqr

import (
	"log"
	"net"
)

// 暂未启用

func DnsSafeCheck(url string) bool {
	cname, err := net.LookupCNAME(url)
	if err != nil {
		return false
	}
	log.Println(cname)
	return false
}
