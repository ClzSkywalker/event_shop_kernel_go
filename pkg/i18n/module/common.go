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
	taskContentCode    = 130001
	classifyModuleCode = 140001
	teamModuleCode     = 150001

	SystemErrorCode = 10001 + iota
	StructToStructErr
	TranslatorNotFoundErr
	RequestParamBindErr
	ReqMissErr
	DbErrorErr
	EncryptPwdErr
	GenerateUlidErr
	TokenInvalid
	TokenExpired
	OperateNoPermission
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: SuccessCode, Msg: ""},
		{Tag: language.English, Key: TaskModeErr, Msg: ""},

		{Tag: language.Chinese, Key: SystemErrorCode, Msg: "系统内部错误:%s"},
		{Tag: language.English, Key: SystemErrorCode, Msg: "Internal system error:%s"},

		{Tag: language.Chinese, Key: StructToStructErr, Msg: "系统类型转换错误:%s"},
		{Tag: language.English, Key: StructToStructErr, Msg: "System type conversion error:%s"},

		{Tag: language.Chinese, Key: TranslatorNotFoundErr, Msg: "校验器: %s 未找到"},
		{Tag: language.English, Key: TranslatorNotFoundErr, Msg: "validator: %s not found"},

		{Tag: language.Chinese, Key: RequestParamBindErr, Msg: "参数传递错误:%s"},
		{Tag: language.English, Key: RequestParamBindErr, Msg: "Parameter passing error:%s"},

		{Tag: language.Chinese, Key: ReqMissErr, Msg: "请求信息缺失"},
		{Tag: language.English, Key: ReqMissErr, Msg: "Request missing information"},

		{Tag: language.Chinese, Key: EncryptPwdErr, Msg: "创建密码错误，请换一个密码"},
		{Tag: language.English, Key: EncryptPwdErr, Msg: "Incorrect password creation, please change the password"},

		{Tag: language.Chinese, Key: TokenInvalid, Msg: "密钥无效"},
		{Tag: language.English, Key: TokenInvalid, Msg: "Invalid key"},

		{Tag: language.Chinese, Key: TokenExpired, Msg: "密钥过期"},
		{Tag: language.English, Key: TokenExpired, Msg: "Key expiration"},

		{Tag: language.Chinese, Key: OperateNoPermission, Msg: "没有操作权限"},
		{Tag: language.English, Key: OperateNoPermission, Msg: "Have no operation permission"},
	}
	entry.SetEntries(entries...)
}
