package db

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type sqliteDbStruct struct {
	db          *gorm.DB
	log         *zap.Logger
	isInit      bool                     // 是否已被初始化
	curVersion  int                      // current version
	lastVersion int                      // last version
	migrateList []constx.AutoMigrateFunc // migrate func list
	CreateFunc  []constx.CreateTableFunc
	DropFunc    []constx.DropTableFunc
	InitData    func() (err error)
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 获取sqlite的版本
 * @return          {*}
 */
func (s *sqliteDbStruct) GetVersion() (err error) {
	var version int
	err = s.db.Raw("pragma user_version").Find(&version).Error
	s.curVersion = version
	if version >= 0 {
		s.isInit = true
	}
	return
}

func (s *sqliteDbStruct) SetVersion() (err error) {
	return s.db.Exec(fmt.Sprintf("pragma user_version=%d", s.lastVersion)).Error
}

func (s *sqliteDbStruct) SetDb(db *gorm.DB) {
	s.db = db
}

func (s *sqliteDbStruct) SetCreateFunc(p ...constx.CreateTableFunc) {
	s.CreateFunc = p
}

func (s *sqliteDbStruct) SetDropFunc(p ...constx.DropTableFunc) {
	s.DropFunc = p
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-07
 * @Description    : 初始化数据库表与数据
 * @param           {string} mode
 * @param           {chan<-DbInitStateType} ch
 * @return          {*}
 */
func (s *sqliteDbStruct) OnInitDb(mode string, ch chan<- constx.DbInitStateType) (err error) {
	if mode == gin.TestMode {
		err = s.onDrop()
		if err != nil {
			return
		}
	}

	ch <- DbCreating
	err = s.onCreate()
	if err != nil {
		return
	}

	err = s.onInitData()
	if err != nil {
		return
	}

	ch <- DbUpgrading
	err = s.onUpgrade()
	return
}

func (s *sqliteDbStruct) onCreate() (err error) {
	if s.curVersion > 0 {
		return
	}
	s.log.Info("init database oncreate start", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	for i := 0; i < len(s.CreateFunc); i++ {
		err = s.CreateFunc[i]()
		if err != nil {
			return
		}
	}
	err = s.SetVersion()
	if err != nil {
		return
	}
	s.curVersion = s.lastVersion
	s.log.Info("init database oncreate end", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	return
}

func (s *sqliteDbStruct) onUpgrade() (err error) {
	s.log.Info("upgrade database onupgrade start", zap.Int(
		"oldVersion", s.curVersion),
		zap.Int("newVersion", s.lastVersion))

	for i := s.curVersion; i < s.lastVersion; i++ {
		f := s.migrateList[i-1]
		err = f()
		if err != nil {
			s.log.Error("upgrade database err",
				zap.Int("upgrate version", i),
				zap.Int(
					"oldVersion", s.curVersion),
				zap.Int("newVersion", s.lastVersion))
			return
		}
	}

	s.curVersion = s.lastVersion
	s.log.Info("upgrade database onupgrade end", zap.Int(
		"oldVersion", s.curVersion),
		zap.Int("newVersion", s.lastVersion))
	return
}

func (s *sqliteDbStruct) onDrop() (err error) {
	s.log.Info("init database ondrop start", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	for i := 0; i < len(s.DropFunc); i++ {
		err = s.DropFunc[i]()
		if err != nil {
			return
		}
	}
	err = s.db.Exec("pragma user_version=0").Error
	if err != nil {
		return
	}
	s.isInit = false
	s.curVersion = 0
	s.log.Info("init database ondrop end", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	return
}

func (s *sqliteDbStruct) onInitData() (err error) {
	if s.isInit {
		return
	}
	err = s.InitData()
	return
}
