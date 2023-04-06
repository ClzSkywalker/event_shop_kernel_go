package test

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func initGormAndVar() (ctx *contextx.Contextx) {
	loggerx.DbLog = zap.NewExample()
	loggerx.ZapLog = zap.NewExample()
	gdb, idb, err := db.InitTestDatabase()
	if err != nil {
		panic(err)
	}
	err = gdb.Transaction(func(tx *gorm.DB) error {
		idb = container.InitIDB(idb, tx)
		ch := make(chan constx.DbInitStateType, 10)
		return idb.OnInitDb(gin.TestMode, ch)
	})
	if err != nil {
		panic(err)
	}
	container.GlobalServerContext = container.NewBaseServiceContext(container.GlobalServerContext, gdb)
	ctx, err = newCtx()
	if err != nil {
		panic(err)
	}
	return ctx
}

func newCtx() (ctx *contextx.Contextx, err error) {
	ctx = &contextx.Contextx{Language: constx.LangChinese}
	ctx.BaseTx = *container.NewBaseServiceContext(container.GlobalServerContext, container.GlobalServerContext.Db)
	_, err = service.UserRegisterByUid(ctx)
	return
}

func refreshDB(ctx *contextx.Contextx) {
	ctx.BaseTx = *container.NewBaseServiceContext(container.GlobalServerContext, container.GlobalServerContext.Db)
}
