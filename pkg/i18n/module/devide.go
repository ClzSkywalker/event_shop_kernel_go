package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	DevideInsertErr = devideModuleCode + iota
	DevideQueryErr
	DevideNotfoundErr
	DevideUpdateErr
	DevideDeleteErr
	DevideTitleRepeatErr
	DevideDelExistTaskErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: DevideQueryErr, Msg: "分组查询失败"},
		{Tag: language.English, Key: DevideQueryErr, Msg: "Grouping query failure"},

		{Tag: language.Chinese, Key: DevideNotfoundErr, Msg: "分组不存在"},
		{Tag: language.English, Key: DevideNotfoundErr, Msg: "Devide does not exist"},

		{Tag: language.Chinese, Key: DevideInsertErr, Msg: "插入分组失败"},
		{Tag: language.English, Key: DevideInsertErr, Msg: "Insert devide failed"},

		{Tag: language.Chinese, Key: DevideUpdateErr, Msg: "更新分组失败"},
		{Tag: language.English, Key: DevideUpdateErr, Msg: "Update devide failed"},

		{Tag: language.Chinese, Key: DevideDeleteErr, Msg: "删除分组失败"},
		{Tag: language.English, Key: DevideDeleteErr, Msg: "Delete devide failed"},

		{Tag: language.Chinese, Key: DevideTitleRepeatErr, Msg: "在同一分类下，存在相同分组名字"},
		{Tag: language.English, Key: DevideTitleRepeatErr, Msg: "Under the same classification, there are the same group name"},

		{Tag: language.Chinese, Key: DevideDelExistTaskErr, Msg: "删除分组失败，该分组存在任务"},
		{Tag: language.English, Key: DevideDelExistTaskErr, Msg: "Delete group fail, the task group exists"},
	}
	entry.SetEntries(entries...)
}
