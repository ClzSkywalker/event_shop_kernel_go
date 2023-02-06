package container

import (
	"flag"

	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
)

const (
	KernelVersion = "v0.0.1"
)

type AppConfig struct {
	Port          int                // 端口
	Mode          string             // gin mode
	KernelVersion string             // 内核版本
	DbPath        string             // sqlite path
	DbInitState   db.DbInitStateType // 数据库是否初始化完毕
	Language      int                // 0-zh 1-en
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 解析参数,初始化config
 * @return          {*}
 */
func InitConfig(port, local int, mode, dbPath string) {
	var port1, local1 int
	var mode1, dbPath1 string
	flag.IntVar(&port1, "port", port, "--port")
	flag.StringVar(&mode1, "mode", mode, "--mode")
	flag.StringVar(&dbPath1, "dbPath", dbPath, "database path")
	flag.IntVar(&local1, "local", local, "language")
	flag.Parse()
	config := AppConfig{
		Port:          port1,
		Mode:          mode1,
		DbPath:        dbPath1,
		Language:      local1,
		KernelVersion: KernelVersion,
	}
	utils.ZapLog.Info(`dbInit`,
		zap.String("db path", config.DbPath),
		zap.Int("local", config.Language))
	GlobalServerContext.Config = config
}
