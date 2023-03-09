package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TeamRepeatErr = teamModuleCode + iota
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TeamRepeatErr, Msg: "团队ID已存在"},
		{Tag: language.English, Key: TeamRepeatErr, Msg: "The team ID already exists"},
	}
	entry.SetEntries(entries...)
}
