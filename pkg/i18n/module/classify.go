package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	ClassifyQueryErr = classifyModuleCode + iota
	ClassifyInsertErr
	ClassifyTitleRepeatErr
	ClassifyUpdateErr
	ClassifyNotfoundErr
	ClassifyDeleteErr
	ClassifyDelExistDevideErr
	ClassifyDeepErr
	ClassifyParentNoExistErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: ClassifyQueryErr, Msg: "查询分类失败"},
		{Tag: language.English, Key: ClassifyQueryErr, Msg: "Class query failure"},

		{Tag: language.Chinese, Key: ClassifyInsertErr, Msg: "分类创建失败"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "Class creation failure"},

		{Tag: language.Chinese, Key: ClassifyTitleRepeatErr, Msg: "分类名字重复"},
		{Tag: language.English, Key: ClassifyTitleRepeatErr, Msg: "Category name repetition"},

		{Tag: language.Chinese, Key: ClassifyUpdateErr, Msg: "分类更新失败"},
		{Tag: language.English, Key: ClassifyUpdateErr, Msg: "Class update failure"},

		{Tag: language.Chinese, Key: ClassifyNotfoundErr, Msg: "分类不存在"},
		{Tag: language.English, Key: ClassifyNotfoundErr, Msg: "Classification does not exist"},

		{Tag: language.Chinese, Key: ClassifyDeleteErr, Msg: "分类删除失败"},
		{Tag: language.English, Key: ClassifyDeleteErr, Msg: "Class deletion failure"},

		{Tag: language.Chinese, Key: ClassifyDelExistDevideErr, Msg: "分类中存在分组，不可删除"},
		{Tag: language.English, Key: ClassifyDelExistDevideErr, Msg: "Classification of grouping, cannot be deleted"},

		{Tag: language.Chinese, Key: ClassifyDeepErr, Msg: "分类层次太深"},
		{Tag: language.English, Key: ClassifyDeepErr, Msg: "The classification level is too deep"},

		{Tag: language.Chinese, Key: ClassifyParentNoExistErr, Msg: "父文件夹不存在"},
		{Tag: language.English, Key: ClassifyParentNoExistErr, Msg: "The parent folder does not exist"},
	}
	entry.SetEntries(entries...)
}
