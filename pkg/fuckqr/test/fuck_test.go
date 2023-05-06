package test

import (
	"maan/pkg/fuckqr"
	"testing"
)

func TestDnsSafeCheck(t *testing.T) {
	_ = fuckqr.DnsSafeCheck("baidu.com")
}
