package utils

import (
	"log"
	"testing"
	"time"
)

func TestRandomString(t *testing.T) {
	// 生成随机文件名
	randomString, err := GenerateRandomString(8)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(randomString)
	now := time.Now()
	timestamp := now.UnixNano()
	log.Println("timestamp:", timestamp)

}
