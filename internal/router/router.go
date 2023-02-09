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
	globalRoute.Handle(http.MethodGet, "/hello", ctrl.GetHello)
	globalRoute.Handle(http.MethodGet, "/pwd", ctrl.GetPwd)
	globalRoute.Handle(http.MethodGet, "/create", ctrl.Create)
}
