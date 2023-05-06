package dict

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestDelword(t *testing.T) {
	data, err := os.ReadFile("default_dict.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if line == "犬" {
			log.Print(i + 1)
			log.Println(line)
		}
	}
}
