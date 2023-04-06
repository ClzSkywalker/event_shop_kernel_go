package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	ClassifyQueryErr = classifyModuleCode + iota
	ClassifyInsertErr
	ClassifyExistedErr
	ClassifyUpdateErr
	ClassifyNotfoundErr
	ClassifyDeleteErr
	ClassifyDelExistTask
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: ClassifyQueryErr, Msg: "查询分类失败"},
		{Tag: language.English, Key: ClassifyQueryErr, Msg: "Class query failure"},
		{Tag: language.Chinese, Key: ClassifyInsertErr, Msg: "分类创建失败"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "Class creation failure"},
		{Tag: language.Chinese, Key: ClassifyInsertErr, Msg: "分类已存在"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "Classification already exists"},
		{Tag: language.Chinese, Key: ClassifyInsertErr, Msg: "分类更新失败"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "Class update failure"},
		{Tag: language.Chinese, Key: ClassifyInsertErr, Msg: "分类不存在"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "Classification does not exist"},
		{Tag: language.Chinese, Key: ClassifyInsertErr, Msg: "分类删除失败"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "Class deletion failure"},
		{Tag: language.Chinese, Key: ClassifyDelExistTask, Msg: "该分类中存任务，不可删除"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "This category contains tasks and cannot be deleted"},
	}
	entry.SetEntries(entries...)
}
