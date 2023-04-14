package test

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
)

const (
	nikeName1 = "skywalker"
	email1    = "123@qq.com"
	// email2    = "1231@qq.com"
	// email3    = "1232@qq.com"
	phone1 = "123456"
	pwd1   = "123456"
)

var (
	classify1 = entity.ClassifyInsertReq{
		Title:     "test111",
		ShowType:  constx.TaskShowNormal,
		OrderType: constx.TaskOrderDefault,
		Color:     "#ffffff",
		Sort:      5,
	}

	classify2 = entity.ClassifyInsertReq{
		Title:     "test222",
		ShowType:  constx.TaskShowNormal,
		OrderType: constx.TaskOrderDefault,
		Color:     "#fffabd",
		Sort:      6,
	}
)
