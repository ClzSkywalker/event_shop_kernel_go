package utils

import "encoding/json"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

func NewResultByBytes(data []byte) (res *Result, err error) {
	err = json.Unmarshal(data, res)
	return
}
