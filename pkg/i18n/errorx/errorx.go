package errorx

type CodeError struct {
	Code  int64  `json:"code"`
	Msg   string `json:"msg"`
	Field []interface{}
}

func (e CodeError) Error() string {
	return e.Msg
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-10
 * @Description    : 返回 error
 * @param           {int64} code
 * @param           {...interface{}} fields 变量参数
 * @return          {*}
 */
func NewCodeError(code int64, fields ...interface{}) (err CodeError) {
	err = CodeError{Code: code, Field: fields}
	return
}
func Is(err error, code int64) bool {
	switch t := err.(type) {
	case CodeError:
		return t.Code == code
	default:
		return false
	}
}
