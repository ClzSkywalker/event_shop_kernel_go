package entity

type DevideInsertReqEntity struct {
	ClassifyId string `json:"classify_id" binding:"required" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Sort       int    `json:"sort" validate:"required"`
}

type DevideUpdateReqEntity struct {
	OnlyCode   string `json:"oc" binding:"required" validate:"required"`
	ClassifyId string `json:"classify_id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Sort       int    `json:"sort" validate:"required"`
}
