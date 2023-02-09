package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TaskModeErr = taskModeModuleCode + iota
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TaskModeErr, Msg: "task mode"},
		{Tag: language.English, Key: TaskModeErr, Msg: "task mode"},
	}
	entry.SetEntries(entries...)
}
