package utils

import (
	"golang.org/x/crypto/scrypt"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 加密
 * @param           {*} salt
 * @param           {string} pwd
 * @return          {*}
 */
func EncryptPwd(pwd, salt string) (epwd string, err error) {
	pwdByte, err := scrypt.Key([]byte(pwd), []byte(salt), 1<<15, 8, 1, 32)
	epwd = string(pwdByte)
	return
}
