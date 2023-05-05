package fuckqr

import (
	"bufio"
	"fmt"
	"log"
	"maan/pkg/dao"
	"maan/pkg/model"
	"os"
	"strings"
	"testing"
)

func TestHackerIp(t *testing.T) {
	file, err := os.Open("ipsum.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	// 计数器
	var count int = 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) > 0 {
			d := dao.NewIpSumDao()
			hackerIp := &model.IpSum{
				HackerIp: fields[0],
			}
			_ = d.SaveHackerIp(nil, nil, hackerIp)
			log.Println(count)
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file:", err)
	}
}
