package validate

import (
	"fmt"
	"testing"
)

func TestCheckUsername(t *testing.T) {
	testName := []string{
		"abcd",
		"风撒大沙发",
		"--_--   ",
		"s",
		" ",
		"",
	}
	for i, s := range testName {
		ok := VerifyUsername(s)
		t.Log("test", i, " :", s, " -> ", ok)
	}
}

func TestCheckPassword(t *testing.T) {
	testPwd := []string{
		"abcd",
		"abcdefg7",
		"qpalzmkl_1",
		"我是撒大苏打阿萨sc-",
		"aaaadsde -1",
		"--__--1s",
		"asdfghjklqwer123_",
		"asdfghjklqwer123_-",
		"asdfghjklqwer1_---",
		"asdfghjklqwer1_",
		"1234567890abcde-",
		"1s------=====+++",
	}
	for i, s := range testPwd {
		ok := VerifyPassword(s)
		fmt.Println("test", i, " :", s, " -> ", ok)
	}
}
