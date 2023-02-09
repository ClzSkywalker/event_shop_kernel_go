package middleware

import (
	"github.com/gin-gonic/gin"
)

func TransMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		byteList := make([]byte, 0, 1<<10)
		_, err := c.Request.Response.Body.Read(byteList)
		if err != nil {
			return
		}
		// res, err := utils.NewResultByBytes(byteList)
		// if err != nil {
		// 	return
		// }
		// i18n.Trans(c.Request.Header.Get("language"), res.Code)
		// res.Code
	}
}
