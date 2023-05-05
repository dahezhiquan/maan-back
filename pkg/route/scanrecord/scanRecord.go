package scanrecord

import (
	"github.com/gin-gonic/gin"
	"log"
	"maan/common"
	"maan/pkg/service"
)

func init() {
	log.Println("init scanRecord route")
	common.Register(&RouterScanRecord{})
}

type RouterScanRecord struct {
}

func (*RouterScanRecord) Route(r *gin.Engine) {
	h := service.NewHandlerScanRecord()
	g := r.Group("/scan")
	g.GET("/info", h.AnalysisQrScan)
}
