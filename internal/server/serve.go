package server

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/middleware"
	"github.com/clz.skywalker/event.shop/kernal/internal/router"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/recoverx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : pc server
 * @return          {*}
 */
func CmdServer() {
	loggerx.ZapLog.Info(`start event shop kernel server`,
		zap.String("version", constx.KernelVersion))
	serveInit()
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : mobile server
 * @param           {*} port
 * @param           {int} local
 * @param           {*} mode
 * @param           {string} dbPath
 * @return          {*}
 */
func KernelServer(c container.AppConfig) {
	err := container.InitConfig(c)
	if err != nil {
		panic(err)
	}
	loggerx.ZapLog.Info(`start event shop kernel server`,
		zap.String("version", constx.KernelVersion))

	serveInit()
}

func serveInit() {
	container.GlobalServerContext.Validator = i18n.NewParaValidation()
	ch := make(chan constx.DbInitStateType, 1)
	go recoverx.RecoverReadChanFunc(container.GlobalServerContext.Config.LogPath, container.TailDbInitStatus, ch)
	go recoverx.RecoverWriteChanFunc(container.GlobalServerContext.Config.LogPath, container.InitServiceContext, ch)
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
	ginServer.Use(gin.CustomRecovery(middleware.RecoveryMiddleware))
	ginServer.Use(middleware.CorsMiddleware())
	ginServer.Use(middleware.LoggerMiddleware())
	router.RouterManager(ginServer)
	err := ginServer.Run(fmt.Sprintf(":%d", container.GlobalServerContext.Config.Port))
	if err != nil {
		loggerx.ZapLog.Error(`server error`,
			zap.Error(err),
			zap.Int("part", container.GlobalServerContext.Config.Port))
	}
}
