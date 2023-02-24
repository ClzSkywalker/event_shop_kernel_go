package container

import (
	"flag"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
)

type AppConfig struct {
	Port          int                    `json:"port"`           // 端口
	Mode          string                 `json:"mode"`           // gin mode
	KernelVersion string                 `json:"kernel_version"` // 内核版本
	DbPath        string                 `json:"db_path"`        // sqlite path
	LogPath       string                 `json:"log_path"`       // db path
	DbInitState   constx.DbInitStateType `json:"db_state"`       // 数据库是否初始化完毕
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 解析参数,初始化config
 * @return          {*}
 */
func InitConfig(c AppConfig) error {
	var port int
	var mode, dbPath, logPath string
	flag.IntVar(&port, "port", c.Port, "--port")
	flag.StringVar(&mode, "mode", c.Mode, "--mode")
	flag.StringVar(&dbPath, "dbPath", c.DbPath, "database path")
	flag.StringVar(&logPath, "logPath", c.LogPath, "log path")
	flag.Parse()
	config := AppConfig{
		Port:          port,
		Mode:          mode,
		DbPath:        dbPath,
		LogPath:       logPath,
		KernelVersion: constx.KernelVersion,
	}
	err := utils.CreateDir(logPath)
	if err != nil {
		return err
	}
	loggerx.ZapLog = loggerx.InitLogger(c.LogPath, "kernel", true)
	loggerx.DbLog = loggerx.InitLogger(c.LogPath, "db", false)
	loggerx.ReqLog = loggerx.InitLogger(c.LogPath, "req", true)
	loggerx.ZapLog.Info(`dbInit`,
		zap.Any("config", config))
	GlobalServerContext.Config = config
	return nil
}
