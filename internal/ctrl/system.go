package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/gin-gonic/gin"
)

/**
 * @Author         : Angular
 * @Date           : 2023-02-10
 * @Description    : 获取kernel配置
 * @param           {*gin.Context} c
 * @return          {*}
 */
func KernelState(c *gin.Context) {
	ret := httpx.NewResult()
	ret.Data = container.GlobalServerContext.Config
	c.JSON(http.StatusOK, ret)
}
