package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		reqData, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqData))
		c.Next()
		endTime := time.Now()
		execTime := endTime.Sub(startTime)
		reqUri := c.Request.RequestURI
		reqPath := c.Request.URL.Path

		// 日志白名单
		if reqUri == "" {
			return
		}

		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		loggerx.ReqLog.Info("request", zap.Int("code", statusCode),
			zap.Duration("exec_time", execTime),
			zap.String("client_ip", clientIP),
			zap.String("req_path", reqPath),
			zap.String("req_params", string(reqData)))
	}
}
