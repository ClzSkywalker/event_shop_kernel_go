package container

import (
	"sync"

	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GlobalServerContext *BaseServiceContext = &BaseServiceContext{}

type BaseServiceContext struct {
	Config           AppConfig
	Db               *gorm.DB
	TaskModel        model.ITaskModel
	TaskChildModel   model.ITaskChildModel
	TaskContentModel model.ITaskContentModel
	TaskModelModel   model.ITaskModeModel
	ClassifyModel    model.IClassifyModel
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : 初始化 context
 * @param           {chan<-db.DbInitStateType} ch
 * @return          {*}
 */
func InitServiceContext(ch chan<- db.DbInitStateType) {
	database, err := db.InitDatabase(GlobalServerContext.Config.DbPath, ch)
	if err != nil {
		utils.ZapLog.Error(`init Database error`,
			zap.Error(err),
			zap.String("version", GlobalServerContext.Config.KernelVersion),
			zap.String("path", GlobalServerContext.Config.DbPath))
		return
	}
	once := new(sync.Once)
	once.Do(func() {
		GlobalServerContext.Db = database
		GlobalServerContext.TaskModel = model.NewDefaultTaskModel(database)
		GlobalServerContext.TaskChildModel = model.NewDefaultTaskChildModel(database)
		GlobalServerContext.TaskContentModel = model.NewDefaultTaskContentModel(database)
		GlobalServerContext.TaskModelModel = model.NewDefaultTaskModeModel(database)
		GlobalServerContext.ClassifyModel = model.NewDefaultClassifyModel(database)
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
func TailDbInitStatus(ch <-chan db.DbInitStateType) {
	for state := range ch {
		GlobalServerContext.Config.DbInitState = state
		if state == db.DbInitSuccess {
			return
		}
	}
}
