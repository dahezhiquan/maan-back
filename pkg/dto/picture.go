package dto

// 查询图片 req

type FindPicReq struct {
	Type string `json:"type" form:"type" validate:"max=30"`
}

// 查询图片 resp

type PictureResp struct {
	PictureList []string `json:"picture_list"`
}
