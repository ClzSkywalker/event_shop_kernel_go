package loggerx

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ZapLog *zap.Logger
	DbLog  *zap.Logger
	ReqLog *zap.Logger
)

/**
 * @Author         : Angular
 * @Date           : 2023-02-07
 * @Description    : 初始化日志
 * @param           {string} path 日志路径
 * @param           {string} logType 日志类型
 * @param           {bool} addCaller  是否增加caller打印
 * @return          {*}
 */
func InitLogger(path, logType string, addCaller bool) *zap.Logger {
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
		zapcore.NewMultiWriteSyncer(getWriterSyncer(path, logType, "info"),
			zapcore.AddSync(os.Stdout)),
		lowPriority)
	errCore := zapcore.NewCore(
		getEncoder(),
		// 双向输出 file,console
		zapcore.NewMultiWriteSyncer(getWriterSyncer(path, logType, "err"),
			zapcore.AddSync(os.Stdout)),
		highPriority)
	if addCaller {
		//zap.AddCaller() 显示文件名 和 行号
		return zap.New(zapcore.NewTee(infoCore, errCore), zap.AddCaller())
	}
	return zap.New(zapcore.NewTee(infoCore, errCore))
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-24
 * @Description    : 日志同步
 * @param           {*} path 路径
 * @param           {*} logType 日志类型
 * @param           {string} level 日志等级
 * @return          {*}
 */
func getWriterSyncer(path, logType, level string) zapcore.WriteSyncer {
	// lsyncer := &lumberjack.Logger{
	// 	Filename:   fmt.Sprintf("%s/%s.log", path, level),
	// 	MaxSize:    10,
	// 	MaxBackups: 3,
	// 	MaxAge:     30,
	// 	//Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	// 	Compress: false,
	// }
	logFilePath := fmt.Sprintf("%s/%s.%s.log", path, logType, level)
	hook, err := rotatelogs.New(
		logFilePath+".%Y%m%d",
		rotatelogs.WithLinkName(logFilePath),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}

	return zapcore.AddSync(hook)
}

// core 三个参数之  Encoder 编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //在日志文件中使用大写字母记录日志级别
	return zapcore.NewJSONEncoder(encoderConfig)
}
