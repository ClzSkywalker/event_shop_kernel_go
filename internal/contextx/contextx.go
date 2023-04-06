package contextx

import (
	"context"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
)

// 自定义context
// 全局参数

type Contextx struct {
	context.Context `json:"-"`
	Language        string                       `json:"language,omitempty"`
	TID             string                       `json:"tid,omitempty"` // team id
	UID             string                       `json:"uid,omitempty"` // user id
	BaseTx          container.BaseServiceContext `json:"-"`
}
