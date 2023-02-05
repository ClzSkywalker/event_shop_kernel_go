package db

import (
	"fmt"

	"github.com/clz.skywalker/event.shop/kernal/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type sqliteDbStruct struct {
	db          *gorm.DB
	isInit      bool              // 是否已被初始化
	curVersion  int               // current version
	lastVersion int               // last version
	migrateList []autoMigrateFunc // migrate func list
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 获取sqlite的版本
 * @return          {*}
 */
func (s *sqliteDbStruct) getSqliteVersion() (err error) {
	var version int
	err = s.db.Raw("pragma user_version").Find(&version).Error
	s.curVersion = version
	if version >= 0 {
		s.isInit = true
	}
	return
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 初始化数据
 * @return          {*}
 */
func (s *sqliteDbStruct) onCreate() (err error) {
	if s.curVersion > 0 {
		return
	}
	logger.ZapLog.Info("init database oncreate start", zap.Int(
		"oldVersion", s.curVersion,
	), zap.Int("newVersion", s.lastVersion))
	err = s.db.Exec(fmt.Sprintf(`
	CREATE TABLE task(
		id INTEGER NOT NULL PRIMARY KEY  AUTOINCREMENT , --
		title TEXT   , --todo事件
		completed_time INTEGER   , --任务完成时间
		give_up_time INTEGER   , --放弃任务时间
		start_time INTEGER   , --开始时间
		end_time INTEGER   , --结束时间
		classify_id INTEGER   , --所属类别
		content_id INTEGER   , --描述文字id
		task_mode_id INTEGER   , --任务模式
		created_time INTEGER   , --创建时间
		updated_time INTEGER   , --更新时间
		deleted_time INTEGER    --删除时间
	)  ; --
	
	
	CREATE TABLE task_content(
		id INTEGER NOT NULL PRIMARY KEY  AUTOINCREMENT , --
		task_id INTEGER   , --
		content TEXT(900)   , --内容
		file_list BLOB   , --存储文件地址数组
		created_time INTEGER   , --创建时间
		updated_time INTEGER   , --更新时间
		deleted_time INTEGER    --删除时间
	)  ; --
	
	
	CREATE TABLE task_child(
		id INTEGER NOT NULL PRIMARY KEY  AUTOINCREMENT , --
		title TEXT   , --todo事件
		parent_id INTEGER   , --父id
		completed_time INTEGER   , --完成时间
		give_up_time INTEGER   , --放弃任务时间
		created_time INTEGER   , --创建时间
		updated_time INTEGER   , --更新时间
		deleted_time INTEGER    --删除时间
	)  ; --
	
	
	CREATE TABLE task_mode(
		id INTEGER NOT NULL PRIMARY KEY  AUTOINCREMENT , --
		mode_id INTEGER NOT NULL  , --重复模式，1-天，2-周，3-月，4-年，5-工作日(周一-周五)，6-法定工作日，7-法定节假日
		config BLOB   , --所选择的天：{“day”:[1,2]}
		created_time INTEGER   , --创建时间
		updated_time INTEGER   , --更新时间
		deleted_time INTEGER    --删除时间
	)  ; --
	
	
	CREATE TABLE classify(
		id INTEGER NOT NULL PRIMARY KEY  AUTOINCREMENT , --
		title TEXT   , --
		color TEXT   , --颜色(#ffffff)
		sort INTEGER   , --排序
		created_time INTEGER   , --创建时间
		updated_time INTEGER   , --更新时间
		deleted_time INTEGER    --删除时间
	)  ; --

	PRAGMA user_version =%d;
`, lastVersion)).Error
	s.curVersion = lastVersion
	logger.ZapLog.Info("init database oncreate end", zap.Int(
		"oldVersion", s.curVersion),
		zap.Int("newVersion", s.lastVersion))
	return
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-05
 * @Description    : 版本升级
 * @return          {*}
 */
func (s *sqliteDbStruct) onUpgrade() (err error) {
	logger.ZapLog.Info("upgrade database onupgrade start", zap.Int(
		"oldVersion", s.curVersion),
		zap.Int("newVersion", s.lastVersion))
	for i := s.curVersion; i < s.lastVersion; i++ {
		f := s.migrateList[i-1]
		err = f()
		if err != nil {
			return
		}
	}
	s.curVersion = s.lastVersion
	logger.ZapLog.Info("upgrade database onupgrade end", zap.Int(
		"oldVersion", s.curVersion),
		zap.Int("newVersion", s.lastVersion))
	return
}
