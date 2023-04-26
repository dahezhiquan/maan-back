package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"maan/common"
	"maan/common/errs"
	"maan/pkg/connections/database/transaction"
	"maan/pkg/domain"
	"maan/pkg/dto"
	"net/http"
)

type HandlerPicture struct {
	pictureDomain *domain.PictureDomain
	tx            *transaction.Transaction
}

func NewHandlerPicture() *HandlerPicture {
	return &HandlerPicture{
		pictureDomain: domain.NewPictureDomain(),
		tx:            transaction.NewTransaction(),
	}
}

func (h *HandlerPicture) FindPicture(ctx *gin.Context) {
	var result = &common.Result{}
	var req = dto.FindPicReq{}
	var resp = dto.PictureResp{}

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

	pics, err := h.pictureDomain.FindPic(req.Type)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(err))
		return
	}

	for _, pic := range pics {
		resp.PictureList = append(resp.PictureList, pic.Src)
	}

	// 防止返回 null 值
	if resp.PictureList == nil {
		resp.PictureList = make([]string, 0)
	}

	ctx.JSON(http.StatusOK, result.Success(resp))
	return
}
