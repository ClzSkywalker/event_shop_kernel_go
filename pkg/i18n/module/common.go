package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	SuccessCode = 0

	// task mode
	userModeuleColde   = 100001
	taskModeModuleCode = 110001
	taskModuleCode     = 120001
	classifyModuleCode = 130001

	SystemErrorCode = 10001 + iota
	TranslatorNotFoundErr
	RequestParamBindErr
	DbErrorErr
	EncryptPwdErr
	GenerateUlidErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: SuccessCode, Msg: ""},
		{Tag: language.English, Key: TaskModeErr, Msg: ""},

		{Tag: language.Chinese, Key: SystemErrorCode, Msg: "系统内部错误:%s"},
		{Tag: language.English, Key: SystemErrorCode, Msg: "Internal system error:%s"},

		{Tag: language.Chinese, Key: TranslatorNotFoundErr, Msg: "校验器: %s 未找到"},
		{Tag: language.English, Key: TranslatorNotFoundErr, Msg: "validator: %s not found"},

		{Tag: language.Chinese, Key: RequestParamBindErr, Msg: "参数传递错误:%s"},
		{Tag: language.English, Key: RequestParamBindErr, Msg: "Parameter passing error:%s"},

		{Tag: language.Chinese, Key: EncryptPwdErr, Msg: "创建密码错误，请换一个密码"},
		{Tag: language.English, Key: EncryptPwdErr, Msg: "Incorrect password creation, please change the password"},
	}
	entry.SetEntries(entries...)
}
