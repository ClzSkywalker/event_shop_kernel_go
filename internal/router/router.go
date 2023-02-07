package router

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func RouterManager(c *gin.Engine) {
	globalRoute := c.Group("api/v1")
	globalRoute.Handle("GET", "/hello", ctrl.GetHello)
	globalRoute.Handle("GET", "/pwd", ctrl.GetPwd)
}
