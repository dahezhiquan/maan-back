package fuckqr

import (
	"crypto/tls"
	"maan/common/validate"
	"net/url"
	"strings"
)

func SslCheck(urlContent string) bool {

	// 创建TLS配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false, // 开启证书校验
	}

	// 得到url主机名
	u, err := url.Parse(urlContent)
	if err != nil {
		return false
	}
	host := u.Hostname()

	// 先判断是否存在ssl证书
	isHttps := strings.HasPrefix(urlContent, "https://")
	if !isHttps {
		return false
	}

	// 判断主机名是不是ip地址，防止无用的ssl连接导致接口变慢
	isIp := validate.VerifyIpFormat(host)
	if isIp {
		return false
	}

	// 建立连接
	conn, err := tls.Dial("tcp", host+":443", tlsConfig)

	if err != nil {
		// 连接失败
		return false
	}
	defer func(conn *tls.Conn) {
		_ = conn.Close()
	}(conn)

	// 验证证书
	if err = conn.VerifyHostname(host); err != nil {
		// 证书无效
		return false
	}

	// 证书有效
	return true
}
