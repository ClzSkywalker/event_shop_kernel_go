package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TeamRepeatErr = teamModuleCode + iota
	TeamFindErr
	TeamNotFound
	TeamUpdateErr
	TeamDeleteErr
	TeamCreateErr
	TeamNameRepeatErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TeamRepeatErr, Msg: "团队ID已存在"},
		{Tag: language.English, Key: TeamRepeatErr, Msg: "The team ID already exists"},

		{Tag: language.Chinese, Key: TeamFindErr, Msg: "团队查询错误"},
		{Tag: language.English, Key: TeamFindErr, Msg: "Team query error"},

		{Tag: language.Chinese, Key: TeamNotFound, Msg: "团队没有找到"},
		{Tag: language.English, Key: TeamNotFound, Msg: "The team did not find"},

		{Tag: language.Chinese, Key: TeamUpdateErr, Msg: "团队更新失败"},
		{Tag: language.English, Key: TeamUpdateErr, Msg: "Team update failure"},

		{Tag: language.Chinese, Key: TeamDeleteErr, Msg: "团队删除失败"},
		{Tag: language.English, Key: TeamDeleteErr, Msg: "Team deletion failure"},

		{Tag: language.Chinese, Key: TeamCreateErr, Msg: "团队创建失败"},
		{Tag: language.English, Key: TeamCreateErr, Msg: "Team creation failure"},

		{Tag: language.Chinese, Key: TeamNameRepeatErr, Msg: "团队名称重复"},
		{Tag: language.English, Key: TeamNameRepeatErr, Msg: "Team name duplication"},
	}
	entry.SetEntries(entries...)
}
