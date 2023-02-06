package server

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/middleware"
	"github.com/clz.skywalker/event.shop/kernal/internal/router"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ServerStart() {
	logger.ZapLog.Info(`start event shop kernel server`,
		zap.String("version", container.KernelVersion))

	container.InitConfig()

	ch := make(chan db.DbInitStateType, 1)
	go container.TailDbInitStatus(ch)
	go container.InitServiceContext(ch)
	go serve()
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : http启动
 * @return          {*}
 */
func serve() {
	if container.GlobalServerContext.Config.Mode != gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	ginServer := gin.New()
	// 兼容插入较大的资源文件时内存占用较大
	ginServer.MaxMultipartMemory = 1024 * 1024 * 32
	ginServer.Use(gin.Recovery())
	ginServer.Use(middleware.CorsMiddleware())
	router.RouterManager(ginServer)
	err := ginServer.Run(fmt.Sprintf(":%d", container.GlobalServerContext.Config.Port))
	if err != nil {
		logger.ZapLog.Error(`server error`,
			zap.Error(err),
			zap.Int("part", container.GlobalServerContext.Config.Port))
	}
}
