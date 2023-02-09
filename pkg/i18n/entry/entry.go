package entry

import "golang.org/x/text/language"

var entries []Entry

type Entry struct {
	Tag language.Tag
	Key int
	Msg interface{}
}

func SetEntries(entry ...Entry) {
	entries = append(entries, entry...)
}

func GetEntries() []Entry {
	return entries
}
