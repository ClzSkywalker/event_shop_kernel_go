package errorx

import (
	"go.uber.org/zap/zapcore"
)

type CodeError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func MewCodeError(Code int64, msg string, fields ...ErrorField) CodeError {
	for i := 0; i < len(fields); i++ {
		switch fields[i].Type {
		case zapcore.StringType:
			msg = fields[i].replaceString(msg, fields[i].String)
		case zapcore.Int32Type:
			msg = fields[i].replaceInt(msg, fields[i].Int)
		case zapcore.Int64Type:
			msg = fields[i].replaceInt64(msg, fields[i].Int64)
		case zapcore.Float32Type:
			msg = fields[i].replaceFloat32(msg, fields[i].Float32)
		case zapcore.Float64Type:
			msg = fields[i].replaceFloat64(msg, fields[i].Float64)
		default:
		}
	}
	return CodeError{Code: Code, Msg: msg}
}
