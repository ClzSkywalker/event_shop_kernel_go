package db

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbGorm *gorm.DB

const (
	lastVersion = 1
)

func InitDatabase(dbPath string) (err error) {
	utils.ZapLog.Info("init database start")
	DbGorm, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return
	}
	var db iinitDb = &sqliteDbStruct{
		lastVersion: lastVersion,
		migrateList: make([]autoMigrateFunc, 0),
	}
	err = db.getSqliteVersion()
	if err != nil {
		return
	}
	err = db.onCreate()
	if err != nil {
		return
	}
	err = db.onUpgrade()
	if err != nil {
		return
	}
	utils.ZapLog.Info("init database end")
	return
}

type autoMigrateFunc func() (err error)

type iinitDb interface {
	getSqliteVersion() (err error)
	onCreate() (err error)
	onUpgrade() (err error)
}
