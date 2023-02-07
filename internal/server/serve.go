package server

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/middleware"
	"github.com/clz.skywalker/event.shop/kernal/internal/router"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
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
	container.InitConfig(container.AppConfig{
		Port:     4319,
		Mode:     gin.ReleaseMode,
		DbPath:   "",
		LogPath:  "./logs",
		Language: 0,
	})

	utils.ZapLog.Info(`start event shop kernel server`,
		zap.String("version", container.KernelVersion))

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
	container.InitConfig(c)

	utils.ZapLog.Info(`start event shop kernel server`,
		zap.String("version", container.KernelVersion))

	serveInit()
}

func serveInit() {
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
		utils.ZapLog.Error(`server error`,
			zap.Error(err),
			zap.Int("part", container.GlobalServerContext.Config.Port))
	}
}
