package module

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"golang.org/x/text/language"
)

const (
	ClassifyQueryErr = classifyModuleCode + iota
)

func init() {
	var entries = []entry.Entry{
		{Tag: language.Chinese, Key: ClassifyQueryErr, Msg: "查询分类失败: %s"},
		{Tag: language.English, Key: ClassifyQueryErr, Msg: "Class query failure: %s"},
	}
	entry.SetEntries(entries...)
}
