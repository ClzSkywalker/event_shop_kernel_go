package utils

import jsoniter "github.com/json-iterator/go"

var JsonUtil = jsoniter.ConfigCompatibleWithStandardLibrary

//
// Author         : ClzSkywalker
// Date           : 2023-04-06
// Description    :
// param           {*} T
// param           {interface{}} V 这个值一定是一个指针
// return          {*}
//
func StructToStruct(T, V interface{}) (err error) {
	byteList, err := JsonUtil.Marshal(T)
	if err != nil {
		return
	}
	err = JsonUtil.Unmarshal(byteList, V)
	return
}
