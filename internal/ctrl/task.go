package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	ret := utils.NewResult()
	ret.Msg = "Hello world!"
	c.JSON(http.StatusOK, ret)
}
