package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 创建一个ulid
 * @return          {*}
 */
func NewUlid() (id string, err error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	u, err := ulid.New(ms, entropy)
	id = u.String()
	return
}
