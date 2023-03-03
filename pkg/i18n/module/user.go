package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	UserRegisterErr = userModeuleColde + iota
	UserRegisterRepeatErr
	UserNotFoundErr
	UserPwdErr
	UserPhoneErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: UserRegisterErr, Msg: "用户注册失败"},
		{Tag: language.English, Key: UserRegisterErr, Msg: "User registration failure"},

		{Tag: language.Chinese, Key: UserRegisterRepeatErr, Msg: "邮箱/电话已被注册"},
		{Tag: language.English, Key: UserRegisterRepeatErr, Msg: "Email/phone number has been registered"},

		{Tag: language.Chinese, Key: UserNotFoundErr, Msg: "用户不存在"},
		{Tag: language.English, Key: UserNotFoundErr, Msg: "User does not exist"},

		{Tag: language.Chinese, Key: UserPwdErr, Msg: "密码错误"},
		{Tag: language.English, Key: UserPwdErr, Msg: "Password error"},

		{Tag: language.Chinese, Key: UserPwdErr, Msg: "号码异常"},
		{Tag: language.English, Key: UserPwdErr, Msg: "Abnormal number"},
	}
	entry.SetEntries(entries...)
}
