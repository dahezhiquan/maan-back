package test

import (
	"log"
	"maan/common/api"
	"testing"
)

func TestGetIpInfo(t *testing.T) {
	info, _ := api.GetIpInfo("140.82.112.3")
	log.Println(info.Adcode.O)
}
