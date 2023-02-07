package db

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type sqliteDbStruct struct {
	db          *gorm.DB
	isInit      bool              // 是否已被初始化
	curVersion  int               // current version
	lastVersion int               // last version
	migrateList []autoMigrateFunc // migrate func list
	CreateFunc  []CreateTableFunc
	DropFunc    []DropTableFunc
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

func (s *sqliteDbStruct) SetCreateFunc(p ...CreateTableFunc) {
	s.CreateFunc = p
}

func (s *sqliteDbStruct) SetDropFunc(p ...DropTableFunc) {
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
func (s *sqliteDbStruct) OnInitDb(mode string, ch chan<- DbInitStateType) (err error) {
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
	utils.ZapLog.Info("init database oncreate start", zap.Int(
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
	utils.ZapLog.Info("init database oncreate end", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	return
}

func (s *sqliteDbStruct) onUpgrade() (err error) {
	utils.ZapLog.Info("upgrade database onupgrade start", zap.Int(
		"oldVersion", s.curVersion),
		zap.Int("newVersion", s.lastVersion))

	for i := s.curVersion; i < s.lastVersion; i++ {
		f := s.migrateList[i-1]
		err = f()
		if err != nil {
			utils.ZapLog.Error("upgrade database err",
				zap.Int("upgrate version", i),
				zap.Int(
					"oldVersion", s.curVersion),
				zap.Int("newVersion", s.lastVersion))
			return
		}
	}

	s.curVersion = s.lastVersion
	utils.ZapLog.Info("upgrade database onupgrade end", zap.Int(
		"oldVersion", s.curVersion),
		zap.Int("newVersion", s.lastVersion))
	return
}

func (s *sqliteDbStruct) onDrop() (err error) {
	utils.ZapLog.Info("init database ondrop start", zap.Int(
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
	utils.ZapLog.Info("init database ondrop end", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	return
}

func (s *sqliteDbStruct) onInitData() (err error) {
	if s.isInit {
		return
	}
	utils.ZapLog.Info("init data start", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	err = s.db.Exec(`
	-- classify
INSERT INTO classify(title,color,sort,created_time,updated_time)  VALUES('普通','#c7ecee',1,strftime('%s','now'),strftime('%s','now'));
INSERT INTO classify(title,color,sort,created_time,updated_time)  VALUES('工作','#4ac8b2',2,strftime('%s','now'),strftime('%s','now'));

-- task_mode
INSERT INTO task_mode(mode_id)  VALUES(0);

-- task
INSERT INTO task(title,classify_id,task_mode_id,created_time,updated_time)  VALUES('欢迎加入',1,1,strftime('%s','now'),strftime('%s','now'));
INSERT INTO task(title,classify_id,task_mode_id,created_time,updated_time)  VALUES('第一天',1,1,strftime('%s','now'),strftime('%s','now'));

-- task_content
INSERT INTO task_content(task_id,content)  VALUES(1,'(｡･∀･)ﾉﾞ嗨，小当家，欢迎回到您的小卖铺！');
INSERT INTO task_content(task_id,content)  VALUES(2,'第一天，您打算做些什么呢？');`).Error
	if err != nil {
		return
	}
	utils.ZapLog.Info("init data end", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	return
}
