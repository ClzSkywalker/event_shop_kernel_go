package httpx

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
)

type Result struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

func (r *Result) SetCodeErr(errx error) {
	switch e := errx.(type) {
	case errorx.CodeError:
		r.Code = e.Code
		r.Msg = e.Error()
	default:
		r.Code = module.SystemErrorCode
		r.Msg = errx.Error()
	}
}
