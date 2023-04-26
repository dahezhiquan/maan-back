package picture

import (
	"github.com/gin-gonic/gin"
	"log"
	"maan/common"
	"maan/pkg/service"
)

func init() {
	log.Println("init picture route")
	common.Register(&RouterPicture{})
}

type RouterPicture struct {
}

func (*RouterPicture) Route(r *gin.Engine) {
	h := service.NewHandlerPicture()
	g := r.Group("/index")
	g.GET("/picture", h.FindPicture)
}
