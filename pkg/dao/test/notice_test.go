package test

import (
	"context"
	"log"
	"maan/pkg/dao"
	"testing"
	"time"
)

// 测试查询公告
func TestFindNotice(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	d := dao.NewNoticeDao()
	notices, _ := d.FindNotice(nil, ctx)
	log.Println("======测试查询公告======")
	for _, notice := range notices {
		log.Println(notice.Title)
	}
}
