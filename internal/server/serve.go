package server

import (
	"flag"

	"github.com/clz.skywalker/event.shop/kernal/internal/config"
	"github.com/clz.skywalker/event.shop/kernal/internal/middleware"
	"github.com/clz.skywalker/event.shop/kernal/internal/router"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Serve() {
	utils.SetupZapLogger()
	utils.ZapLog.Info(`start event shop kernel server`,
		zap.String("version", config.AppConfig.KernelVersion))

	flagInit()

	err := db.InitDatabase(config.AppConfig.DbPath)
	if err != nil {
		utils.ZapLog.Error(`init Database error`,
			zap.Error(err),
			zap.String("version", config.AppConfig.KernelVersion),
			zap.String("path", config.AppConfig.DbPath))
		return
	}

	// gin.SetMode(gin.ReleaseMode)
	ginServer := gin.New()
	// 兼容插入较大的资源文件时内存占用较大
	ginServer.MaxMultipartMemory = 1024 * 1024 * 32
	ginServer.Use(gin.Recovery())
	ginServer.Use(middleware.CorsMiddleware())
	router.RouterManager(ginServer)
	err = ginServer.Run(":4935")
	if err != nil {
		utils.ZapLog.Error(`server error`,
			zap.Error(err),
			zap.String("part", ":4935"))
	}
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 解析参数
 * @return          {*}
 */
func flagInit() {
	dbPath := flag.String("dbPath", "", "database path")
	local := flag.Int("local", 0, "language")
	flag.Parse()
	config.AppConfig.DbPath = *dbPath
	config.AppConfig.Language = *local
	utils.ZapLog.Info(`dbInit`,
		zap.String("db path", config.AppConfig.DbPath),
		zap.Int("local", config.AppConfig.Language))
}
