package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TaskInsertErr = taskModuleCode + iota
	TaskUpdateErr
	TaskDeleteErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TaskInsertErr, Msg: "插入任务失败"},
		{Tag: language.English, Key: TaskInsertErr, Msg: "Insert task failed"},

		{Tag: language.Chinese, Key: TaskInsertErr, Msg: "更新任务失败"},
		{Tag: language.English, Key: TaskInsertErr, Msg: "Update task failed"},

		{Tag: language.Chinese, Key: TaskInsertErr, Msg: "删除任务失败"},
		{Tag: language.English, Key: TaskInsertErr, Msg: "Delete task failed"},
	}
	entry.SetEntries(entries...)
}
