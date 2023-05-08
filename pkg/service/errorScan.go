package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"maan/common"
	"maan/common/errs"
	"maan/pkg/connections/database/transaction"
	"maan/pkg/domain"
	"maan/pkg/dto"
	"maan/pkg/model"
	"net/http"
)

type HandlerErrorScan struct {
	errorScanDomain *domain.ErrorScanDomain
	tx              *transaction.Transaction
}

func NewHandlerErrorScan() *HandlerErrorScan {
	return &HandlerErrorScan{
		errorScanDomain: domain.NewErrorScanDomain(),
		tx:              transaction.NewTransaction(),
	}
}

func (h *HandlerErrorScan) SaveErrorScan(ctx *gin.Context) {
	var result = &common.Result{}
	var req = dto.ErrorScanReq{}

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

	errorScan := &model.ErrorScan{
		Content:      req.Content,
		Mvss:         req.Mvss,
		ContentTitle: req.ContentTitle,
		IsDfa:        req.IsDfa,
		RiskType:     req.RiskType,
		IpAddr:       req.IpAddr,
	}

	err := h.errorScanDomain.SaveErrorScan(errorScan)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		return
	}

	ctx.JSON(http.StatusOK, result.Success("成功"))
	return
}
