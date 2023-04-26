package service

import (
	"github.com/gin-gonic/gin"
	"maan/common"
	"maan/pkg/connections/database/transaction"
	"maan/pkg/domain"
	"maan/pkg/dto"
	"net/http"
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

	// 防止返回 null 值
	if resp.NoticeList == nil {
		resp.NoticeList = make([]string, 0)
	}

	ctx.JSON(http.StatusOK, result.Success(resp))
	return
}
