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
	UserBindNoUidLoginErr
	UserBindedEmailErr
	UserBindedPhoneErr
	UserEmailBindByOtherErr
	UserPhoneBindByOtherErr
	UserUpdateErr
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

		{Tag: language.Chinese, Key: UserBindNoUidLoginErr, Msg: "用户已绑定其他登录方式，不支持uid登录"},
		{Tag: language.English, Key: UserBindNoUidLoginErr, Msg: "The user has been bound to another login method and does not support uid login"},

		{Tag: language.Chinese, Key: UserBindedEmailErr, Msg: "用户已绑定邮箱"},
		{Tag: language.English, Key: UserBindedEmailErr, Msg: "The user has bound a email"},

		{Tag: language.Chinese, Key: UserBindedPhoneErr, Msg: "用户已绑定电话"},
		{Tag: language.English, Key: UserBindedPhoneErr, Msg: "The user has a bound phone"},

		{Tag: language.Chinese, Key: UserBindedEmailErr, Msg: "邮箱已被其他用户绑定"},
		{Tag: language.English, Key: UserBindedEmailErr, Msg: "The mailbox has been bound to another user"},

		{Tag: language.Chinese, Key: UserBindedEmailErr, Msg: "电话已被其他用户绑定"},
		{Tag: language.English, Key: UserBindedEmailErr, Msg: "The phone has been bound to another user"},

		{Tag: language.Chinese, Key: UserBindedEmailErr, Msg: "用户更新失败"},
		{Tag: language.English, Key: UserBindedEmailErr, Msg: "User update failure"},
	}
	entry.SetEntries(entries...)
}
