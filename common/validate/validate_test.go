package validate

import (
	"log"
	"testing"
)

func TestUrlCheck(t *testing.T) {
	isUrl := VerifyUrlFormat("https://baidu.com/a/php")
	log.Println(isUrl)
}
