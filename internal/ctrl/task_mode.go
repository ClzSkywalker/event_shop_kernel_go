package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	ret := utils.NewResult()
	// tm := model.TaskModeModel{}
	// err := c.ShouldBindJSON(&tm)
	// service.InsertTaskMode(container.GlobalServerContext.TaskModelModel)
	msg := i18n.Trans(c.GetHeader("language"), errorx.MewCodeError(module.TaskModeErr,
		"abc", 10))
	ret.Msg = msg
	c.JSON(http.StatusOK, ret)
}
