package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"secureQR/common"
	"secureQR/pkg/connections/database/transaction"
	"secureQR/pkg/domain"
	"secureQR/pkg/dto"
)

type HandlerNotice struct {
	noticeDomain *domain.NoticeDomain
	tx           *transaction.Transaction
}

func NewHandlerNotice() *HandlerNotice {
	return &HandlerNotice{
		noticeDomain: domain.NewNoticeDomain(),
		tx:           transaction.NewTransaction(),
	}
}

func (h *HandlerNotice) FindNotice(ctx *gin.Context) {
	var result = &common.Result{}
	var resp = dto.NoticeResp{}

	// 获取公告内容
	notices, err := h.noticeDomain.FindNotice()
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		return
	}

	// 将notices的值 移植 到resp中
	for _, notice := range notices {
		resp.NoticeList = append(resp.NoticeList, notice.Title)
	}

	ctx.JSON(http.StatusOK, result.Success(resp))
	return
}
