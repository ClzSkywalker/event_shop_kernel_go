package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ZapLog *zap.Logger
)

/**
 * @Author         : Angular
 * @Date           : 2023-02-07
 * @Description    : 初始化日志
 * @param           {string} path 日志路径
 * @return          {*}
 */
func InitLogger(path string) {
	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})
	infoCore := zapcore.NewCore(
		getEncoder(),
		// 双向输出 file,console
		zapcore.NewMultiWriteSyncer(getWriterSyncer(path, "info"),
			zapcore.AddSync(os.Stdout)),
		lowPriority)
	errCore := zapcore.NewCore(
		getEncoder(),
		// 双向输出 file,console
		zapcore.NewMultiWriteSyncer(getWriterSyncer(path, "err"),
			zapcore.AddSync(os.Stdout)),
		highPriority)
	//zap.AddCaller() 显示文件名 和 行号
	ZapLog = zap.New(zapcore.NewTee(infoCore, errCore), zap.AddCaller())
}

func getWriterSyncer(path, level string) zapcore.WriteSyncer {
	lsyncer := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", path, level),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		//Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
		Compress: false,
	}
	return zapcore.AddSync(lsyncer)
}

// core 三个参数之  Encoder 编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //在日志文件中使用大写字母记录日志级别
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func NewDbLog(log *zap.Logger, config logger.Config, mode string) logger.Interface {
	return &dbLog{
		Log:    log,
		Config: config,
		Mode:   mode,
	}
}

type dbLog struct {
	Mode string
	logger.Config
	Log *zap.Logger
}

func (db *dbLog) LogMode(level logger.LogLevel) logger.Interface {
	db.LogLevel = level
	return db
}
func (db *dbLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if db.LogLevel < logger.Info {
		return
	}
	db.Log.Info("db info:", zap.Any("data", data))
}

func (db *dbLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if db.LogLevel < logger.Warn {
		return
	}
	db.Log.Warn("db ware:"+msg, zap.Any("data", data))
}

func (db *dbLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if db.LogLevel < logger.Error {
		return
	}
	db.Log.Error("db err:"+msg, zap.Any("data", data))
}

func (db *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if db.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && db.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !db.IgnoreRecordNotFoundError):
		sql, rows := fc()
		db.Log.Error("db trace err:", zap.Error(err), zap.Float64("ms", float64(elapsed.Nanoseconds())/1e6), zap.Int64("rows", rows), zap.String("sql", sql))
	case elapsed > db.SlowThreshold && db.SlowThreshold != 0 && db.LogLevel >= logger.Warn:
		sql, rows := fc()
		db.Log.Info("db trace ware:", zap.Duration("SLOW SQL>=", db.SlowThreshold), zap.Int64("rows", rows), zap.String("sql", sql))
	case db.LogLevel == logger.Info:
		sql, rows := fc()
		db.Log.Info("db trace info:", zap.Float64("ms", float64(elapsed.Nanoseconds())/1e6), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}
