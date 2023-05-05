package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"maan/common"
	"maan/common/api"
	"maan/common/errs"
	"maan/common/tms"
	validate2 "maan/common/validate"
	"maan/pkg/connections/database/transaction"
	"maan/pkg/dto"
	"net/http"
)

type HandlerScanRecord struct {
	tx *transaction.Transaction
}

func NewHandlerScanRecord() *HandlerScanRecord {
	return &HandlerScanRecord{
		tx: transaction.NewTransaction(),
	}
}

func (h *HandlerScanRecord) AnalysisQrScan(ctx *gin.Context) {
	var result = &common.Result{}
	var req = dto.QrScanReq{}
	var resp = dto.QrScanResp{}

	// 参数绑定
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(errs.ParameterError))
		return
	}

	// 参数校验
	validate := validator.New()
	err2 := validate.Struct(req)
	if err2 != nil {
		ctx.JSON(http.StatusOK, result.Fail(errs.ParameterError))
		return
	}

	// 判断扫描结果是否是一个url类型的地址，如果是，则进行地址风险检测，否则对内容进行检测
	if !validate2.VerifyUrlFormat(req.Content) {
		// TODO 内容风险检测算法
		ctx.JSON(http.StatusOK, result.Success(resp))
		return
	}

	ip := tms.UrlToIp(req.Content)

	// 得到ip归属地信息
	ipInfo, err := api.GetIpInfo(ip)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
	}

	resp.Mvss = 0
	resp.IpAddr = ipInfo.Adcode.O
	resp.RiskType = "未知类型"

	ctx.JSON(http.StatusOK, result.Success(resp))
	return
}
