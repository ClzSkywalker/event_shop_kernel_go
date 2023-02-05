package server

import (
	"flag"
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/middleware"
	"github.com/clz.skywalker/event.shop/kernal/internal/router"
	"github.com/clz.skywalker/event.shop/kernal/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Serve() {
	logger.SetupZapLogger()
	logger.ZapLog.Info(`start event shop kernel server`,
		zap.String("version", container.KernelVersion))

	config := initFlag()
	err := container.InitServiceContext(config)
	if err != nil {
		logger.ZapLog.Error(`init Database error`,
			zap.Error(err),
			zap.String("path", config.DbPath))
		return
	}

	if config.Mode != gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	ginServer := gin.New()
	// 兼容插入较大的资源文件时内存占用较大
	ginServer.MaxMultipartMemory = 1024 * 1024 * 32
	ginServer.Use(gin.Recovery())
	ginServer.Use(middleware.CorsMiddleware())
	router.RouterManager(ginServer)
	err = ginServer.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logger.ZapLog.Error(`server error`,
			zap.Error(err),
			zap.Int("part", config.Port))
	}
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 解析参数
 * @return          {*}
 */
func initFlag() container.AppConfig {
	port := flag.Int("port", 4935, "port")
	mode := flag.String("mode", "debug", "gin mode")
	dbPath := flag.String("dbPath", "", "database path")
	local := flag.Int("local", 0, "language")
	flag.Parse()
	config := container.AppConfig{
		Port:     *port,
		Mode:     *mode,
		DbPath:   *dbPath,
		Language: *local,
	}
	logger.ZapLog.Info(`dbInit`,
		zap.String("db path", config.DbPath),
		zap.Int("local", config.Language))
	return config
}
