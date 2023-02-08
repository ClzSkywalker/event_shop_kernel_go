package container

import (
	"flag"

	"github.com/clz.skywalker/event.shop/kernal/pkg/consts"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
)

const (
	KernelVersion = "v0.0.1"
)

type AppConfig struct {
	Port          int                    // 端口
	Mode          string                 // gin mode
	KernelVersion string                 // 内核版本
	Language      int                    // 0-zh 1-en
	DbPath        string                 // sqlite path
	LogPath       string                 // db path
	DbInitState   consts.DbInitStateType // 数据库是否初始化完毕
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 解析参数,初始化config
 * @return          {*}
 */
func InitConfig(c AppConfig) {
	var port, local int
	var mode, dbPath, logPath string
	flag.IntVar(&port, "port", c.Port, "--port")
	flag.StringVar(&mode, "mode", c.Mode, "--mode")
	flag.IntVar(&local, "local", c.Language, "language")
	flag.StringVar(&dbPath, "dbPath", c.DbPath, "database path")
	flag.StringVar(&logPath, "logPath", c.LogPath, "log path")
	flag.Parse()
	config := AppConfig{
		Port:          port,
		Mode:          mode,
		Language:      local,
		DbPath:        dbPath,
		LogPath:       logPath,
		KernelVersion: KernelVersion,
	}
	utils.InitLogger(c.LogPath)
	utils.ZapLog.Info(`dbInit`,
		zap.Any("config", config))

	GlobalServerContext.Config = config
}
