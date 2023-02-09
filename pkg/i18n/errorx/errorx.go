package errorx

import "fmt"

type CodeError struct {
	Code  int64  `json:"code"`
	Msg   string `json:"msg"`
	Field []interface{}
}

func (e CodeError) Error() string {
	return fmt.Sprintf(e.Msg, e.Field...)
}

func MewCodeError(code int64, fields ...interface{}) CodeError {
	return CodeError{Code: code, Field: fields}
}
