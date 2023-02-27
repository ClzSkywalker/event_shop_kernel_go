package entity

type ClassifyInsertReq struct {
	Title string `json:"title"`
	Color string `json:"color"`
	Sort  int    `json:"sort"`
}
