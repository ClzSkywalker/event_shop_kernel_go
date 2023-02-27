package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	ClassifyQueryErr = classifyModuleCode + iota
	ClassifyInsertErr
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: ClassifyQueryErr, Msg: "查询分类失败"},
		{Tag: language.English, Key: ClassifyQueryErr, Msg: "Class query failure"},
		{Tag: language.Chinese, Key: ClassifyInsertErr, Msg: "分类创建失败"},
		{Tag: language.English, Key: ClassifyInsertErr, Msg: "Class creation failure"},
	}
	entry.SetEntries(entries...)
}
