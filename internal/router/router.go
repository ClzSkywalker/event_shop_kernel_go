package router

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/ctrl"
	"github.com/gin-gonic/gin"
)

/**
 * @Author         : Angular
 * @Date           : 2023-02-07
 * @Description    : 路由管理
 * @param           {*gin.Engine} c
 * @return          {*}
 */
func RouterManager(c *gin.Engine) {
	globalRoute := c.Group("api/v1")
	kernel := globalRoute.Group("/kernel")
	kernel.Handle(http.MethodGet, "/config", ctrl.KernelState)

	globalRoute.Handle(http.MethodPost, "/classify", ctrl.InsertClassify)

	globalRoute.Handle(http.MethodPost, "/task", ctrl.InsertTask)
	globalRoute.Handle(http.MethodPost, "/task_mode", ctrl.CreateTaskMode)
}
