package test

import (
	"log"
	"maan/pkg/domain"
	"testing"
)

// 测试获取图片
func TestFindPicture(t *testing.T) {
	d := domain.NewPictureDomain()
	pics, _ := d.FindPic("indexswiper")
	log.Println("======测试获取图片======")
	for _, pic := range pics {
		log.Println(pic.Src)
	}
}
