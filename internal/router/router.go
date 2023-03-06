package router

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/ctrl"
	"github.com/clz.skywalker/event.shop/kernal/internal/middleware"
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

	userUnAuth := globalRoute.Group("/user/unauth")
	{
		userUnAuth.Handle(http.MethodPost, "/register/email", ctrl.RegisterByEmail)
		userUnAuth.Handle(http.MethodPost, "/register/phone", ctrl.RegisterByPhone)
		userUnAuth.Handle(http.MethodPost, "/register/uid", ctrl.RegisterByUid)

		userUnAuth.Handle(http.MethodPost, "/login/email", ctrl.LoginByEmail)
		userUnAuth.Handle(http.MethodPost, "/login/phone", ctrl.LoginByPhone)
		userUnAuth.Handle(http.MethodPost, "/login/uid", ctrl.LoginByUid)
	}

	userAuth := globalRoute.Group("/user/auth").Use(middleware.JwtMiddleware())
	{
		userAuth.Handle(http.MethodPost, "/bind/email", ctrl.BindEmailByUid)
		userAuth.Handle(http.MethodPost, "/bind/phone", ctrl.BindPhoneByUid)
	}

	classify := globalRoute.Group("/classify")
	{
		classify.Handle(http.MethodPost, "/classify", ctrl.InsertClassify)
	}
	globalRoute.Handle(http.MethodPost, "/task", ctrl.InsertTask)
	globalRoute.Handle(http.MethodPost, "/task_mode", ctrl.CreateTaskMode)
}
