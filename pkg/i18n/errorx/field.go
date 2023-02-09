package errorx

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// error field

func NewErrFieldString(val string) ErrorField {
	return ErrorField{Type: zapcore.StringType, String: val}
}

func NewErrFieldInt(val int) ErrorField {
	return ErrorField{Type: zapcore.Int32Type, Int: val}
}

func NewErrFieldInt64(val int64) ErrorField {
	return ErrorField{Type: zapcore.Int64Type, Int64: val}
}

func NewErrFieldFloat32(val float32) ErrorField {
	return ErrorField{Type: zapcore.Float32Type, Float32: val}
}

func NewErrFieldFloat64(val float64) ErrorField {
	return ErrorField{Type: zapcore.Float64Type, Float64: val}
}

type ErrorField struct {
	Type    zapcore.FieldType
	String  string
	Int     int
	Int64   int64
	Float32 float32
	Float64 float64
}

func (e ErrorField) replaceString(msg string, val string) string {
	return fmt.Sprintf(msg, val)
}

func (e ErrorField) replaceInt(msg string, val int) string {
	return fmt.Sprintf(msg, val)
}
func (e ErrorField) replaceInt64(msg string, val int64) string {
	return fmt.Sprintf(msg, val)
}

func (e ErrorField) replaceFloat32(msg string, val float32) string {
	return fmt.Sprintf(msg, val)
}

func (e ErrorField) replaceFloat64(msg string, val float64) string {
	return fmt.Sprintf(msg, val)
}
