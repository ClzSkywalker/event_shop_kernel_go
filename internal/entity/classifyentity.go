package entity

import "github.com/clz.skywalker/event.shop/kernal/pkg/constx"

type ClassifyInsertReq struct {
	Title     string               `json:"title,omitempty" validate:"required"`
	ShowType  constx.TaskShowType  `json:"show_type,omitempty" validate:"required"`
	OrderType constx.TaskOrderType `json:"order_type,omitempty" validate:"required"`
	Color     string               `json:"color,omitempty" validate:"required"`
	Sort      int                  `json:"sort,omitempty" validate:"required"`
	ParentId  string               `json:"parent_id,omitempty"`
}

type ClassifyDelReq struct {
	OnlyCode string `json:"oc" binding:"required" validate:"required"`
}

type ClassifyItem struct {
	OnlyCode  string               `json:"oc,omitempty"`
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
		OnlyCode string `json:"oc" validate:"required"`
		Sort     int    `json:"sort" validate:"required"`
	}
}

type ClassifyUpdateReq struct {
	OnlyCode  string               `json:"oc,omitempty" binding:"required" validate:"required"`
	Title     string               `json:"title,omitempty" validate:"required"`
	ShowType  constx.TaskShowType  `json:"show_type,omitempty" validate:"required"`
	OrderType constx.TaskOrderType `json:"order_type,omitempty" validate:"required"`
	Color     string               `json:"color,omitempty" validate:"required"`
	Sort      int                  `json:"sort,omitempty" validate:"required"`
	ParentId  string               `json:"parent_id,omitempty"`
}
