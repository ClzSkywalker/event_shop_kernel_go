package router

import (
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
	globalRoute.Handle("GET", "/hello", ctrl.GetHello)
	globalRoute.Handle("GET", "/pwd", ctrl.GetPwd)
}
