package utils

import (
	"testing"
	"time"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/golang-jwt/jwt/v5"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJwtToken(t *testing.T) {
	Convey("test token", t, func() {
		token := ""
		var err error
		Convey("generate token", func() {
			claim := GenerateClaims(jwt.RegisteredClaims{
				Issuer:    constx.TokenIssuer,
				ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(constx.TokenExpiresAt * time.Second)},
				Subject:   constx.TokenSub,
				ID:        "t-id",
			}, map[string]interface{}{
				constx.TokenUID: "u-id",
			})
			token, err = GenerateToken(constx.TokenSecret, claim)
			So(err, ShouldEqual, nil)
			Convey("parse token", func() {
				jtoken, err := ParseToken(token, constx.TokenSecret)
				So(err, ShouldEqual, nil)
				So(jtoken.Claims.(jwt.MapClaims)[constx.TokenUID], ShouldEqual, "u-id")
			})
		})
	})
}
