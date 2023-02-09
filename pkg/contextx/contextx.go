package contextx

import "context"

// 自定义context
// 全局参数

type Contextx struct {
	context.Context
	Language string `json:"language"`
}
