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

var (
	GlobalServerContext *BaseServiceContext = &BaseServiceContext{}
	once                                    = new(sync.Once)
)

type BaseServiceContext struct {
	Config           AppConfig
	Db               *gorm.DB
	Validator        *i18n.ParaValidation
	UserModel        model.IUserModel
	TeamModel        model.ITeamModel
	UserToTeamModel  model.IUserToTeamModel
	ClassifyModel    model.IClassifyModel
	DevideModel      model.IDevideModel
	TaskModel        model.ITaskModel
	TaskContentModel model.ITaskContentModel
	TaskModeModel    model.ITaskModeModel
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
	once.Do(func() {
		GlobalServerContext = NewBaseServiceContext(GlobalServerContext, database)
		err = database.Transaction(func(tx *gorm.DB) (err error) {
			idb = InitIDB(idb, tx)
			err = idb.OnInitDb(GlobalServerContext.Config.Mode, ch)
			if err != nil {
				return err
			}

			// test mode 初始化一个用户
			if GlobalServerContext.Config.Mode == gin.TestMode {
				uid := utils.NewUlid()
				tid := utils.NewUlid()
				err = InitData(tx, constx.LangChinese, uid, tid)
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

func NewBaseServiceContext(base *BaseServiceContext, database *gorm.DB) *BaseServiceContext {
	base.Db = database
	base.UserModel = model.NewDefaultUserModel(database)
	base.TeamModel = model.NewDefaultTeamModel(database)
	base.UserToTeamModel = model.NewDefaultUserToTeamModel(database)
	base.ClassifyModel = model.NewDefaultClassifyModel(database)
	base.DevideModel = model.NewDefaultDevideModel(database)
	base.TaskModel = model.NewDefaultTaskModel(database)
	base.TaskContentModel = model.NewDefaultTaskContentModel(database)
	base.TaskModeModel = model.NewDefaultTaskModeModel(database)
	return base
}

func InitIDB(idb db.IOriginDb, tx *gorm.DB) db.IOriginDb {
	idb.SetDb(tx)
	idb.SetDropFunc(
		model.NewDefaultUserModel(tx).DropTable,
		model.NewDefaultTeamModel(tx).DropTable,
		model.NewDefaultUserToTeamModel(tx).DropTable,
		model.NewDefaultClassifyModel(tx).DropTable,
		model.NewDefaultDevideModel(tx).DropTable,
		model.NewDefaultTaskModel(tx).DropTable,
		model.NewDefaultTaskContentModel(tx).DropTable,
		model.NewDefaultTaskModeModel(tx).DropTable,
	)
	idb.SetCreateFunc(
		model.NewDefaultUserModel(tx).CreateTable,
		model.NewDefaultTeamModel(tx).CreateTable,
		model.NewDefaultUserToTeamModel(tx).CreateTable,
		model.NewDefaultClassifyModel(tx).CreateTable,
		model.NewDefaultDevideModel(tx).CreateTable,
		model.NewDefaultTaskModel(tx).CreateTable,
		model.NewDefaultTaskContentModel(tx).CreateTable,
		model.NewDefaultTaskModeModel(tx).CreateTable,
	)
	return idb
}

//
// Author         : ClzSkywalker
// Date           : 2023-04-06
// Description    : 初始化用户数据
// param           {*gorm.DB} tx
// param           {*} lang
// param           {*} uid
// param           {string} tid
// return          {*}
//
func InitData(tx *gorm.DB, lang, uid, tid string) (err error) {
	defer func() {
		if err != nil {
			err = i18n.NewCodeError(module.UserDataInit)
		}
	}()
	err = model.NewDefaultUserModel(tx).InitData(lang, uid)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	err = model.NewDefaultTeamModel(tx).InitData(lang, uid, tid)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	// 绑定用户入口
	err = model.NewDefaultUserModel(tx).Update(model.UserModel{CreatedBy: uid, TeamIdPort: tid})
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	err = model.NewDefaultUserToTeamModel(tx).InitData(uid, tid)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	cid := utils.NewUlid()
	err = model.NewDefaultClassifyModel(tx).InitData(lang, uid, tid, cid)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	tmid := utils.NewUlid()
	err = model.NewDefaultTaskModeModel(tx).InitData(tmid, tid)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	contentId, err := model.NewDefaultTaskContentModel(tx).InitData()
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	did := utils.NewUlid()
	err = model.NewDefaultDevideModel(tx).InitData(lang, uid, cid, did)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	err = model.NewDefaultTaskModel(tx).InitData(lang, uid, did, tid, contentId)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}

	return
}
