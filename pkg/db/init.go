package db

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	lastVersion = 1
)

func InitDatabase(dbPath string) (db *gorm.DB, err error) {
	logger.ZapLog.Info("init database start")
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return
	}
	var s iinitDb = &sqliteDbStruct{
		db:          db,
		lastVersion: lastVersion,
		migrateList: make([]autoMigrateFunc, 0),
	}
	err = s.getSqliteVersion()
	if err != nil {
		return
	}
	err = s.onCreate()
	if err != nil {
		return
	}
	err = s.onUpgrade()
	if err != nil {
		return
	}
	logger.ZapLog.Info("init database end")
	return
}

type autoMigrateFunc func() (err error)

type iinitDb interface {
	getSqliteVersion() (err error)
	onCreate() (err error)
	onUpgrade() (err error)
}
