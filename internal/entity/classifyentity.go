package entity

import "github.com/clz.skywalker/event.shop/kernal/pkg/constx"

type ClassifyInsertReq struct {
	Title string `json:"title"`
	Color string `json:"color"`
	Sort  int    `json:"sort"`
}

type ClassifyDelReq struct {
	OnlyCode string `json:"only_code" binding:"required"`
}

type ClassifyItem struct {
	OnlyCode  string               `json:"only_code,omitempty"`
	CreatedBy string               `json:"created_by,omitempty"`
	Title     string               `json:"title,omitempty"`
	ShowType  constx.TaskShowType  `json:"show_type,omitempty"`
	OrderType constx.TaskOrderType `json:"order_type,omitempty"`
	Color     string               `json:"color,omitempty"`
	Sort      int                  `json:"sort,omitempty"`
	ParentId  string               `json:"parent_id,omitempty"`
}

type ClassifyFind struct {
	Data []ClassifyItem
}

type ClassifyOrderReq struct {
	Data []struct {
		OnlyCode string `json:"only_code" binding:"required"`
		Sort     int    `json:"sort"`
	}
}
