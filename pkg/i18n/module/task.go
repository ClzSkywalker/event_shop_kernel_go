package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TaskInsertErr = taskModuleCode + iota
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TaskInsertErr, Msg: "插入任务失败: %s"},
		{Tag: language.English, Key: TaskInsertErr, Msg: "Insert task failed: %s"},
	}
	entry.SetEntries(entries...)
}
