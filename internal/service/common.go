package service

import (
	"time"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-09
 * @Description    : 创建token
 * @param           {TokenInfo} t
 * @return          {*}
 */
func GenerateToken(t entity.TokenInfo) (token string, err error) {
	id := utils.NewUlid()
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		return
	}
	claim := utils.GenerateClaims(jwt.RegisteredClaims{
		Issuer:    constx.TokenIssuer,
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(constx.TokenExpiresAt * time.Second)},
		Subject:   constx.TokenSub,
		ID:        id,
	}, map[string]interface{}{
		constx.TokenUID: t.UID,
		constx.TokenTID: t.TID,
	})
	token, err = utils.GenerateToken(constx.TokenSecret, claim)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("claim", claim))
		return
	}
	return
}

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-09
 * @Description    : 解析 token
 * @param           {string} token
 * @return          {*}
 */
func ParseToken(token string) (tokenInfo entity.TokenInfo, err error) {
	t, err := utils.ParseToken(token, constx.TokenSecret)
	if err != nil {
		return
	}
	tmap := t.Claims.(jwt.MapClaims)
	tokenInfo.JMap = tmap
	tokenInfo.UID = tmap[constx.TokenUID].(string)
	tokenInfo.TID = tmap[constx.TokenTID].(string)
	return
}
