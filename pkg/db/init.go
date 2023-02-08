package db

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/pkg/consts"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	lastVersion                          = 1
	DbInitFailure consts.DbInitStateType = iota - 1
	DbCreating
	DbUpgrading
	DbInitSuccess
)

func InitDatabase(dbPath, mode string) (db *gorm.DB, idb iinitDb, err error) {
	utils.ZapLog.Info("init database start")
	dbLog := utils.NewDbLog(utils.ZapLog, logger.Config{
		SlowThreshold:             100,         // 慢 SQL 阈值
		Colorful:                  true,        // 禁用彩色打印
		IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
		LogLevel:                  logger.Info, // 日志级别
	}, mode)
	// https://www.sqlite.org/c3ref/busy_timeout.html
	// _busy_timeout int sqlite3_busy_timeout(sqlite3*, int ms);
	db, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s?_busy_timeout=9999999", dbPath)), &gorm.Config{
		// 表结构初始化的时候不添加外键索引
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   dbLog,
	})
	if err != nil {
		return
	}
	idb = &sqliteDbStruct{
		db:          db,
		lastVersion: lastVersion,
		migrateList: make([]consts.AutoMigrateFunc, 0),
	}

	err = idb.GetVersion()
	if err != nil {
		return
	}
	return
}

type iinitDb interface {
	GetVersion() (err error)
	SetVersion() (err error)
	SetDb(*gorm.DB)
	SetCreateFunc(...consts.CreateTableFunc)
	SetDropFunc(...consts.DropTableFunc)
	OnInitDb(mode string, ch chan<- consts.DbInitStateType) (err error)
	onCreate() (err error)
	onUpgrade() (err error)
	onInitData() (err error)
	onDrop() (err error)
}
