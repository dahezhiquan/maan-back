package dto

// 查询图片 req

type FindPicReq struct {
	Type string `json:"type" form:"type" validate:"max=30"`
}

// 图片

type PictureInfo struct {
	Src string `json:"value"`
}

// 查询图片 resp

type PictureResp struct {
	PictureList []PictureInfo `json:"picture_list"`
}
