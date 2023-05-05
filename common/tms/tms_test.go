package tms

import (
	"log"
	"testing"
)

func TestUrlToIp(t *testing.T) {
	urlString := "https://baidu.com"
	ip := UrlToIp(urlString)
	log.Println(ip)
}
