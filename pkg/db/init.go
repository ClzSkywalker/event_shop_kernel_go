package db

import (
	"fmt"
	"time"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	lastVersion                          = 1
	DbInitFailure constx.DbInitStateType = iota - 1
	DbCreating
	DbUpgrading
	DbInitSuccess
)

func InitDatabase(dbPath, mode string) (db *gorm.DB, idb IOriginDb, err error) {
	loggerx.DbLog.Info("init database start")
	dbLog := loggerx.NewDbLog(loggerx.DbLog, logger.Config{
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

	if mode == gin.DebugMode {
		db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		loggerx.DbLog.Error("connect db error %+v", zap.Error(err))
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	idb = &sqliteDbStruct{
		db:          db,
		log:         loggerx.DbLog,
		lastVersion: lastVersion,
		migrateList: make([]constx.AutoMigrateFunc, 0),
	}

	err = idb.GetVersion()
	if err != nil {
		return
	}
	return
}

func InitTestDatabase() (db *gorm.DB, idb IOriginDb, err error) {
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	idb = &sqliteDbStruct{
		db:          db,
		log:         loggerx.DbLog,
		lastVersion: lastVersion,
		migrateList: make([]constx.AutoMigrateFunc, 0),
	}

	err = idb.GetVersion()
	if err != nil {
		return
	}
	return
}
