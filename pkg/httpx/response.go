package httpx

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
)

type Result struct {
	Lang string      `json:"-"`
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
		e.Msg = i18n.Trans(r.Lang, e)
		r.Code = e.Code
		r.Msg = e.Error()
	default:
		e1 := i18n.NewCodeError(module.SystemErrorCode)
		e1.Msg = i18n.Trans(r.Lang, e1)
		r.Code = e1.Code
		r.Msg = errx.Error()
	}
}
