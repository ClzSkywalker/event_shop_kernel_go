package test

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
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
	container.GlobalServerContext = container.NewBaskServiceContext(container.GlobalServerContext, gdb)
	ctx = &contextx.Contextx{Language: constx.LangChinese}
	um := &model.UserModel{
		CreatedBy: utils.NewUlid(),
	}
	_, err = model.NewDefaultUserModel(gdb).Insert(um)
	if err != nil {
		return
	}
	tid := utils.NewUlid()
	err = container.InitData(gdb, constx.LangChinese, um.CreatedBy, tid)
	if err != nil {
		return
	}
	ctx.TID = um.CreatedBy
	return
}
