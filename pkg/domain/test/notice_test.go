package test

import (
	"log"
	"maan/pkg/domain"
	"testing"
)

// 测试获取最近公告
func TestFindNotice(t *testing.T) {
	d := domain.NewNoticeDomain()
	notices, _ := d.FindNotice()
	log.Println("======测试获取最近公告======")
	for _, notice := range notices {
		log.Println(notice.Title)
	}
}
