package entity

type ClassifyInsertReq struct {
	Title string `json:"title"`
	Color string `json:"color"`
	Sort  int    `json:"sort"`
}

type ClassifyDelReq struct {
	OnlyCode string `json:"only_code" binding:"required"`
}

type ClassifyItem struct {
	OnlyCode  string `json:"only_code"`
	CreatedBy string `json:"created_by"`
	Title     string `json:"title"`
	Color     string `json:"color"`
	Sort      int    `json:"sort"`
	ParentId  string `json:"parent_id"`
}

type ClassifyFind struct {
	Data []ClassifyItem
}
