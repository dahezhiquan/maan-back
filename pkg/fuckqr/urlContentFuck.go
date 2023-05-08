package fuckqr

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"maan/common/dfa"
	"maan/pkg/public"
	"net/http"
	"net/url"
)

// 解析目标的html文档内容，进行风险分析
// 返回待扣除的mvss分数值 - url 标题 - 是否命中了dfa检测

func UrlContentCheck(urlContent string) (subMvss int, urlTitle string, isPassDfa bool) {

	// 解析url
	resp, err := http.Get(urlContent)
	if err != nil {
		return subMvss, "无", false
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// 得到html文档的内容
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return subMvss, "无", false
	}

	// 得到文档标题，尝试dfa命中
	docTitle := doc.Find("title").Text()
	urlTitle = docTitle
	isDfa := dfa.CheckWordByDFA(docTitle)
	if isDfa {
		log.Println("命中dfa")
		isPassDfa = true
		subMvss += public.UnSafeDocContentMvss
	}

	// 检查文档中是否存在表单元素
	if doc.Find("form").Length() > 0 {
		subMvss += public.UnSafeFormMvss
	}

	// 判断是否存在内联或远程脚本的🚩
	var hasInlineJs = false
	var hasInsecureJs = false

	// 检查文档中是否存在可疑的脚本
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("src", "") == "" {
			hasInlineJs = true
		} else {
			scriptURL, err := url.Parse(s.AttrOr("src", ""))
			if err != nil {
				subMvss += 0
			}
			if scriptURL.Scheme != "https" {
				hasInsecureJs = true
			}
		}
	})

	if hasInlineJs {
		subMvss += public.UnSafeInlineScriptMvss
	}
	if hasInsecureJs {
		subMvss += public.UnSafeInsecureScriptMvss
	}

	return subMvss, urlTitle, isPassDfa
}
