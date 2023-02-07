package ctrl

import (
	"net/http"
	"os"

	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	ret := utils.NewResult()
	ret.Msg = "Hello world!"
	c.JSON(http.StatusOK, ret)
}

func GetPwd(c *gin.Context) {
	dir, err := os.Getwd()
	ret := utils.NewResult()
	if err != nil {
		ret.Msg = err.Error()
	} else {
		ret.Msg = dir
	}
	c.JSON(http.StatusOK, ret)
}
