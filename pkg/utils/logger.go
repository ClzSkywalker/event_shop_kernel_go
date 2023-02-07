package utils

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
