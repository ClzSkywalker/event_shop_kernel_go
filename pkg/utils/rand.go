package utils

import (
	"github.com/oklog/ulid/v2"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 创建一个ulid
 * @return          {*}
 */
func NewUlid() (id string) {
	u := ulid.Make()
	id = u.String()
	return
}
