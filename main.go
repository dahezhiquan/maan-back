package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"secureQR/common"
	_ "secureQR/pkg/route"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	common.InitRouter(r)
	err := r.Run("127.0.0.1:8080")
	log.Println(err)
}
