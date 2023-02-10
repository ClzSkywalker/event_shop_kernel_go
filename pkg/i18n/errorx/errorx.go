package errorx

type CodeError struct {
	Code  int64  `json:"code"`
	Msg   string `json:"msg"`
	Field []interface{}
}

func (e CodeError) Error() string {
	return e.Msg
}
