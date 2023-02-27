package container

import (
	"sync"

	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GlobalServerContext *BaseServiceContext = &BaseServiceContext{}

type BaseServiceContext struct {
	Config           AppConfig
	Db               *gorm.DB
	Validator        *i18n.ParaValidation
	UserModel        model.IUserModel
	TaskModel        model.ITaskModel
	TaskChildModel   model.ITaskChildModel
	TaskContentModel model.ITaskContentModel
	TaskModeModel    model.ITaskModeModel
	ClassifyModel    model.IClassifyModel
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : 初始化 context
 * @param           {chan<-db.DbInitStateType} ch
 * @return          {*}
 */
func InitServiceContext(ch chan<- constx.DbInitStateType) {
	database, idb, err := db.InitDatabase(GlobalServerContext.Config.DbPath, GlobalServerContext.Config.Mode)
	if err != nil {
		loggerx.ZapLog.Error(`init Database error`,
			zap.Error(err),
			zap.String("version", GlobalServerContext.Config.KernelVersion),
			zap.String("path", GlobalServerContext.Config.DbPath))
		return
	}
	once := new(sync.Once)
	once.Do(func() {
		GlobalServerContext.Db = database
		GlobalServerContext.UserModel = model.NewDefaultUserModel(database)
		GlobalServerContext.TaskModel = model.NewDefaultTaskModel(database)
		GlobalServerContext.TaskChildModel = model.NewDefaultTaskChildModel(database)
		GlobalServerContext.TaskContentModel = model.NewDefaultTaskContentModel(database)
		GlobalServerContext.TaskModeModel = model.NewDefaultTaskModeModel(database)
		GlobalServerContext.ClassifyModel = model.NewDefaultClassifyModel(database)
		err = database.Transaction(func(tx *gorm.DB) error {
			idb.SetDb(tx)
			idb.SetDropFunc(
				model.NewDefaultUserModel(tx).DropTable,
				model.NewDefaultTaskModel(tx).DropTable,
				model.NewDefaultTaskChildModel(tx).DropTable,
				model.NewDefaultTaskContentModel(tx).DropTable,
				model.NewDefaultTaskModeModel(tx).DropTable,
				model.NewDefaultClassifyModel(tx).DropTable)
			idb.SetCreateFunc(
				model.NewDefaultUserModel(tx).CreateTable,
				model.NewDefaultTaskModel(tx).CreateTable,
				model.NewDefaultTaskChildModel(tx).CreateTable,
				model.NewDefaultTaskContentModel(tx).CreateTable,
				model.NewDefaultTaskModeModel(tx).CreateTable,
				model.NewDefaultClassifyModel(tx).CreateTable)
			return idb.OnInitDb(GlobalServerContext.Config.Mode, ch)
		})
		if err != nil {
			ch <- db.DbInitFailure
			loggerx.ZapLog.Error(`init Database error`,
				zap.Error(err),
				zap.String("version", GlobalServerContext.Config.KernelVersion),
				zap.String("path", GlobalServerContext.Config.DbPath))
			return
		}
		ch <- db.DbInitSuccess
	})
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : 追踪 db 初始化状态
 * @param           {<-chandb.DbInitStateType} ch
 * @return          {*}
 */
func TailDbInitStatus(ch <-chan constx.DbInitStateType) {
	for state := range ch {
		GlobalServerContext.Config.DbInitState = state
	}
}

func NewBaskServiceContext(database *gorm.DB) *BaseServiceContext {
	base := *GlobalServerContext
	base.Db = database
	base.TaskModel = model.NewDefaultTaskModel(database)
	base.TaskChildModel = model.NewDefaultTaskChildModel(database)
	base.TaskContentModel = model.NewDefaultTaskContentModel(database)
	base.TaskModeModel = model.NewDefaultTaskModeModel(database)
	base.ClassifyModel = model.NewDefaultClassifyModel(database)
	return &base
}
