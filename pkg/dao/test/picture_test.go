package test

import (
	"context"
	"log"
	"maan/pkg/dao"
	"testing"
	"time"
)

// 测试查询图片
func TestFindPic(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	d := dao.NewPictureDao()
	pics, _ := d.FindPic(nil, ctx, "indexswiper")
	log.Println("======测试查询图片======")
	for _, pic := range pics {
		log.Println(pic.Src)
	}
}
