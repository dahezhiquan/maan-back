package fuckqr

import (
	"log"
	"maan/common/dfa"
	"maan/pkg/public"
)

// 扫描结果不是一个url，则进行内容解析
// 非url的二维码内容安全性一般比较高

func ContentCheck(content string) (subMvss int, isPassDfa bool) {

	isDfa := dfa.CheckWordByDFA(content)
	if isDfa {
		log.Println("命中dfa")
		isPassDfa = true
		subMvss += public.UnSafeContentMvss
	}

	return subMvss, isPassDfa
}
