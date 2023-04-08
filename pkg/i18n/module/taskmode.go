package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TaskModeErr = taskModeModuleCode + iota
	TaskModeQueryErr
	TaskModeNotFoundErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TaskModeErr, Msg: "插入自定义任务类型失败: %s"},
		{Tag: language.English, Key: TaskModeErr, Msg: "Description Failed to insert a custom task type: %s"},

		{Tag: language.Chinese, Key: TaskModeQueryErr, Msg: "任务类型查询失败"},
		{Tag: language.English, Key: TaskModeQueryErr, Msg: "Task type queries to fail"},

		{Tag: language.Chinese, Key: TaskModeNotFoundErr, Msg: "任务类型不存在"},
		{Tag: language.English, Key: TaskModeNotFoundErr, Msg: "Task type does not exist"},
	}
	entry.SetEntries(entries...)
}
