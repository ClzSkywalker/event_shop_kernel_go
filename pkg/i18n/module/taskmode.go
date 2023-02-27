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
		{Tag: language.Chinese, Key: TaskModeErr, Msg: "插入自定义任务类型失败: %s"},
		{Tag: language.English, Key: TaskModeErr, Msg: "Description Failed to insert a custom task type: %s"},
	}
	entry.SetEntries(entries...)
}
