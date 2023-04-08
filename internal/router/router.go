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

	unAuth := globalRoute.Group("/unauth")
	auth := globalRoute.Group("/auth", middleware.JwtMiddleware())

	userUnAuth := unAuth.Group("/user")
	{
		userUnAuth.Handle(http.MethodPost, "/register/email", ctrl.RegisterByEmail)
		userUnAuth.Handle(http.MethodPost, "/register/phone", ctrl.RegisterByPhone)
		userUnAuth.Handle(http.MethodPost, "/register/uid", ctrl.RegisterByUid)

		userUnAuth.Handle(http.MethodPost, "/login/email", ctrl.LoginByEmail)
		userUnAuth.Handle(http.MethodPost, "/login/phone", ctrl.LoginByPhone)
		userUnAuth.Handle(http.MethodPost, "/login/uid", ctrl.LoginByUid)
	}

	userAuth := auth.Group("/user")
	{
		userAuth.Handle(http.MethodGet, "/userinfo", ctrl.UserGetInfo)
		userAuth.Handle(http.MethodPut, "/userinfo", ctrl.UserUpdate)
		userAuth.Handle(http.MethodPost, "/bind/email", ctrl.BindEmailByUid)
		userAuth.Handle(http.MethodPost, "/bind/phone", ctrl.BindPhoneByUid)
	}

	team := auth.Group("/team")
	{
		team.Handle(http.MethodPost, "", ctrl.TeamCreate)
		team.Handle(http.MethodPut, "", ctrl.TeamUpdate)
		team.Handle(http.MethodDelete, "/:team_id", ctrl.TeamDel)
		team.Handle(http.MethodGet, "", ctrl.TeamFindMyTeam)
	}

	classify := auth.Group("/classify")
	{
		classify.Handle(http.MethodGet, "", ctrl.ClassifyQueryTeam)
		classify.Handle(http.MethodPost, "", ctrl.ClassifyInsert)
		classify.Handle(http.MethodPut, "", ctrl.ClassifyUpdate)
		classify.Handle(http.MethodDelete, "/:only_code", ctrl.ClassifyDel)
	}

	task := auth.Group("/task")
	{
		task.Handle(http.MethodGet, "/:classify_id", ctrl.TaskFindByClassifyId)
		task.Handle(http.MethodPost, "", ctrl.TaskInsert)
		task.Handle(http.MethodPut, "/:only_code", ctrl.TaskUpdate)
		task.Handle(http.MethodDelete, "/:only_code", ctrl.TaskDelete)
	}

	taskMode := auth.Group("/task_mode")
	{
		taskMode.Handle(http.MethodPost, "/task_mode", ctrl.CreateTaskMode)
	}
}
