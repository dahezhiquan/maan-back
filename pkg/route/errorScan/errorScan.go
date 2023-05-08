package errorScan

import (
	"github.com/gin-gonic/gin"
	"log"
	"maan/common"
	"maan/pkg/service"
)

func init() {
	log.Println("init errorScan route")
	common.Register(&RouterErrorScan{})
}

type RouterErrorScan struct {
}

func (*RouterErrorScan) Route(r *gin.Engine) {
	h := service.NewHandlerErrorScan()
	g := r.Group("")
	g.POST("/errorscan", h.SaveErrorScan)
}
