package container

import (
	"flag"

	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/logger"
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
func InitConfig() {
	port := flag.Int("port", 4935, "port")
	mode := flag.String("mode", "debug", "gin mode")
	dbPath := flag.String("dbPath", "", "database path")
	local := flag.Int("local", 0, "language")
	flag.Parse()
	config := AppConfig{
		Port:          *port,
		Mode:          *mode,
		DbPath:        *dbPath,
		Language:      *local,
		KernelVersion: KernelVersion,
	}
	logger.ZapLog.Info(`dbInit`,
		zap.String("db path", config.DbPath),
		zap.Int("local", config.Language))
	GlobalServerContext.Config = config
}
