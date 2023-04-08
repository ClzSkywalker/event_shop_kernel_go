package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TaskContentModeErr = taskContentCode + iota
	TaskContentInsertErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TaskContentInsertErr, Msg: "插入任务内容失败"},
		{Tag: language.English, Key: TaskContentInsertErr, Msg: "Fail to insert task content"},
	}
	entry.SetEntries(entries...)
}
