package notice

import (
	"github.com/gin-gonic/gin"
	"log"
	"secureQR/common"
	"secureQR/pkg/service"
)

func init() {
	log.Println("init notice route")
	common.Register(&RouterNotice{})
}

type RouterNotice struct {
}

func (*RouterNotice) Route(r *gin.Engine) {
	h := service.NewHandlerNotice()
	g := r.Group("/index")
	g.GET("/notice", h.FindNotice)
}
