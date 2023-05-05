package test

import (
	"context"
	"log"
	"maan/pkg/dao"
	"maan/pkg/model"
	"testing"
	"time"
)

func TestSaveHackerIp(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	d := dao.NewIpSumDao()
	hackerIp := &model.IpSum{
		HackerIp: "123.123.123.123",
	}
	_ = d.SaveHackerIp(nil, ctx, hackerIp)
}

func TestFindIpSumByIp(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	d := dao.NewIpSumDao()
	ipSum, _ := d.FindIpSumByIp(nil, ctx, "123.123.123.123")
	log.Println(ipSum.HackerIp)
}
