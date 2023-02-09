package errorx

import (
	"strconv"

	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/entry"
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
func Trans(lang string, key int64, err error) string {
	tag := language.MustParse(lang)
	var p = message.NewPrinter(tag)
	if err != nil {
		return p.Sprintf(key, err.Error())
	}
	return p.Sprintf(key)
}
