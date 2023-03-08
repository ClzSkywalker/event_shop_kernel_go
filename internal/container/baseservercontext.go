package container

import (
	"sync"

	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/db"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var GlobalServerContext *BaseServiceContext = &BaseServiceContext{}

type BaseServiceContext struct {
	Config           AppConfig
	Db               *gorm.DB
	Validator        *i18n.ParaValidation
	UserModel        model.IUserModel
	TeamModel        model.ITeamModel
	UserToTeamModel  model.IUserToTeamModel
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
		GlobalServerContext = NewBaskServiceContext(GlobalServerContext, database)
		err = database.Transaction(func(tx *gorm.DB) (err error) {
			idb = InitIDB(idb, tx)
			err = idb.OnInitDb(GlobalServerContext.Config.Mode, ch)
			if err != nil {
				return err
			}
			if GlobalServerContext.Config.Mode == gin.TestMode {
				um := &model.UserModel{
					CreatedBy: utils.NewUlid(),
				}
				_, err = model.NewDefaultUserModel(tx).Insert(um)
				if err != nil {
					return
				}
				err = InitData(tx, constx.LangChinese, um.CreatedBy)
			}
			return
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

func NewBaskServiceContext(base *BaseServiceContext, database *gorm.DB) *BaseServiceContext {
	base.Db = database
	base.UserModel = model.NewDefaultUserModel(database)
	base.TeamModel = model.NewDefaultTeamModel(database)
	base.UserToTeamModel = model.NewDefaultUserToTeamModel(database)
	base.TaskModel = model.NewDefaultTaskModel(database)
	base.TaskChildModel = model.NewDefaultTaskChildModel(database)
	base.TaskContentModel = model.NewDefaultTaskContentModel(database)
	base.TaskModeModel = model.NewDefaultTaskModeModel(database)
	base.ClassifyModel = model.NewDefaultClassifyModel(database)
	return base
}

func InitIDB(idb db.IOriginDb, tx *gorm.DB) db.IOriginDb {
	idb.SetDb(tx)
	idb.SetDropFunc(
		model.NewDefaultUserModel(tx).DropTable,
		model.NewDefaultTeamModel(tx).DropTable,
		model.NewDefaultUserToTeamModel(tx).DropTable,
		model.NewDefaultTaskModel(tx).DropTable,
		model.NewDefaultTaskChildModel(tx).DropTable,
		model.NewDefaultTaskContentModel(tx).DropTable,
		model.NewDefaultTaskModeModel(tx).DropTable,
		model.NewDefaultClassifyModel(tx).DropTable)
	idb.SetCreateFunc(
		model.NewDefaultUserModel(tx).CreateTable,
		model.NewDefaultTeamModel(tx).CreateTable,
		model.NewDefaultUserToTeamModel(tx).CreateTable,
		model.NewDefaultTaskModel(tx).CreateTable,
		model.NewDefaultTaskChildModel(tx).CreateTable,
		model.NewDefaultTaskContentModel(tx).CreateTable,
		model.NewDefaultTaskModeModel(tx).CreateTable,
		model.NewDefaultClassifyModel(tx).CreateTable)
	return idb
}

func InitData(tx *gorm.DB, lang, uid string) (err error) {
	defer func() {
		if err != nil {
			loggerx.ZapLog.Error(err.Error())
			err = i18n.NewCodeError(module.UserDataInit)
		}
	}()
	err = model.NewDefaultUserModel(tx).InitData(lang, uid)
	if err != nil {
		return
	}

	tid := utils.NewUlid()
	err = model.NewDefaultTeamModel(tx).InitData(lang, uid, tid)
	if err != nil {
		return
	}
	err = model.NewDefaultUserToTeamModel(tx).InitData(uid, tid)
	if err != nil {
		return
	}

	cid := utils.NewUlid()
	err = model.NewDefaultClassifyModel(tx).InitData(lang, uid, tid, cid)
	if err != nil {
		return
	}

	tmid := utils.NewUlid()
	err = model.NewDefaultTaskModeModel(tx).InitData(tmid, tid)
	if err != nil {
		return
	}

	err = model.NewDefaultTaskContentModel(tx).InitData()
	if err != nil {
		return
	}

	err = model.NewDefaultTaskModel(tx).InitData(lang, uid, cid, tid)
	if err != nil {
		return
	}

	err = model.NewDefaultTaskChildModel(tx).InitData()
	if err != nil {
		return
	}
	return
}
