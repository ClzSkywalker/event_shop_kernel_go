package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	TaskInsertErr = taskModuleCode + iota
	TaskQueryErr
	TaskNotfoundErr
	TaskUpdateErr
	TaskDeleteErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: TaskQueryErr, Msg: "任务查询失败"},
		{Tag: language.English, Key: TaskQueryErr, Msg: "Task query failure"},

		{Tag: language.Chinese, Key: TaskNotfoundErr, Msg: "任务不存在"},
		{Tag: language.English, Key: TaskNotfoundErr, Msg: "Task does not exist"},

		{Tag: language.Chinese, Key: TaskInsertErr, Msg: "插入任务失败"},
		{Tag: language.English, Key: TaskInsertErr, Msg: "Insert task failed"},

		{Tag: language.Chinese, Key: TaskUpdateErr, Msg: "更新任务失败"},
		{Tag: language.English, Key: TaskUpdateErr, Msg: "Update task failed"},

		{Tag: language.Chinese, Key: TaskDeleteErr, Msg: "删除任务失败"},
		{Tag: language.English, Key: TaskDeleteErr, Msg: "Delete task failed"},
	}
	entry.SetEntries(entries...)
}
