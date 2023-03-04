package test

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func initGormAndVar() {
	loggerx.DbLog = zap.NewExample()
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
	container.GlobalServerContext = container.NewBaskServiceContext(container.GlobalServerContext, gdb)
}
