package utils

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 创建token
 * @param           {string} sceretKey
 * @param           {jwt.MapClaims} data
 * @return          {*}
 */
func GenerateToken(sceretKey string, data jwt.MapClaims) (token string, err error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err = newToken.SignedString([]byte(sceretKey))
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-02-28
 * @Description    : 创造claim
 * @param           {jwt.RegisteredClaims} claims
 * @param           {map[string]interface{}} data
 * @return          {*}
 */
func GenerateClaims(claims jwt.RegisteredClaims, data map[string]interface{}) (mc jwt.MapClaims) {
	mc = make(jwt.MapClaims)
	mc["iss"] = claims.Issuer
	mc["exp"] = claims.ExpiresAt
	mc["sub"] = claims.Subject
	mc["aud"] = claims.Audience
	mc["iat"] = claims.IssuedAt
	mc["iss"] = "ClzSkywalker"
	mc["jti"] = claims.ID
	for k, v := range data {
		mc[k] = v
	}
	return
}
