package i18n

import (
	"strconv"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	_ "github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

// 翻译模块

func init() {
	entries := entry.GetEntries()
	for _, e := range entries {
		switch msg := e.Msg.(type) {
		case string:
			_ = message.SetString(e.Tag, strconv.Itoa(e.Key), msg)
		case catalog.Message:
			_ = message.Set(e.Tag, strconv.Itoa(e.Key), msg)
		case []catalog.Message:
			_ = message.Set(e.Tag, strconv.Itoa(e.Key), msg...)
		}
	}
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-09
 * @Description    : 翻译
 * @param           {string} lang
 * @param           {int64} key
 * @param           {error} err
 * @return          {*}
 */
func Trans(lang string, err errorx.CodeError) string {
	var tag language.Tag
	switch lang {
	case constx.LangChinese, constx.LangEnglish:
		tag = language.MustParse(lang)
	default:
		tag = language.MustParse(constx.LangEnglish)
	}
	var p = message.NewPrinter(tag)
	return p.Sprintf(strconv.FormatInt(err.Code, 10), err.Field...)
}
