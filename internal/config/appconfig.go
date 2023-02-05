package config

var AppConfig *AppConfigStruct = &AppConfigStruct{
	KernelVersion: "v0.0.1",
}

type AppConfigStruct struct {
	KernelVersion string // 内核版本
	DbPath        string // sqlite path
	Language      int    // 0-zh 1-en
}
