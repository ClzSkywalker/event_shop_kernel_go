package container

import (
	"sync"

	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GlobalServerContext *BaseServiceContext

type BaseServiceContext struct {
	Config           AppConfig
	Db               *gorm.DB
	TaskModel        model.ITaskModel
	TaskChildModel   model.ITaskChildModel
	TaskContentModel model.ITaskContentModel
	TaskModelModel   model.ITaskModeModel
	ClassifyModel    model.IClassifyModel
}

func InitServiceContext(c AppConfig) (err error) {
	db, err := db.InitDatabase(c.DbPath)
	if err != nil {
		logger.ZapLog.Error(`init Database error`,
			zap.Error(err),
			zap.String("version", c.KernelVersion),
			zap.String("path", c.DbPath))
		return
	}
	once := new(sync.Once)
	once.Do(func() {
		tmp := &BaseServiceContext{
			Config:           c,
			Db:               db,
			TaskModel:        model.NewDefaultTaskModel(db),
			TaskChildModel:   model.NewDefaultTaskChildModel(db),
			TaskContentModel: model.NewDefaultTaskContentModel(db),
			TaskModelModel:   model.NewDefaultTaskModeModel(db),
			ClassifyModel:    model.NewDefaultClassifyModel(db),
		}
		GlobalServerContext = tmp
	})
	return
}
