package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	SuccessCode     = 0
	SystemErrorCode = 10001 + iota
	TranslatorNotFoundCode
	RequestParamBindCode
	DbErrorCode

	// task mode
	taskModeModuleCode = 100001
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: SuccessCode, Msg: ""},
		{Tag: language.English, Key: TaskModeErr, Msg: ""},

		{Tag: language.Chinese, Key: SystemErrorCode, Msg: "系统内部错误:%s"},
		{Tag: language.English, Key: SystemErrorCode, Msg: "Internal system error:%s"},

		{Tag: language.Chinese, Key: TranslatorNotFoundCode, Msg: "翻译器%s未找到"},
		{Tag: language.English, Key: TranslatorNotFoundCode, Msg: "translator %s not found"},

		{Tag: language.Chinese, Key: RequestParamBindCode, Msg: "参数传递错误:%s"},
		{Tag: language.English, Key: RequestParamBindCode, Msg: "Parameter passing error:%s"},
	}
	entry.SetEntries(entries...)
}
