package container

const (
	KernelVersion = "v0.0.1"
)

type AppConfig struct {
	Port          int    // 端口
	Mode          string // gin mode
	KernelVersion string // 内核版本
	DbPath        string // sqlite path
	Language      int    // 0-zh 1-en
}
