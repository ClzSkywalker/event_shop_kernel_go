package test

import (
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegister(t *testing.T) {
	initGormAndVar()
	Convey("user register", t, func() {
		Convey("register by emial", func() {
			remailreq := entity.RegisterByEmailReq{
				Email: email1,
			}
			remailreq.NickName = nikeName1
			remailreq.Pwd = pwd1
			uid, err := service.RegisterByEmail(container.GlobalServerContext.UserModel, remailreq)
			So(err, ShouldBeNil)
			So(uid, ShouldNotBeBlank)
			_, err = service.RegisterByEmail(container.GlobalServerContext.UserModel, remailreq)
			So(err, ShouldNotBeNil)
		})

		Convey("register by phone", func() {
			req := entity.RegisterByPhoneReq{Phone: phone1}
			req.NickName = nikeName1
			req.Pwd = pwd1
			uid, err := service.RegisterByPhone(container.GlobalServerContext.UserModel, req)
			So(err, ShouldBeNil)
			So(uid, ShouldNotBeBlank)
			_, err = service.RegisterByPhone(container.GlobalServerContext.UserModel, req)
			So(err, ShouldNotBeNil)
		})

		Convey("register by uid", func() {
			uid, err := service.RegisterByUid(container.GlobalServerContext.UserModel)
			So(err, ShouldBeNil)
			So(uid, ShouldNotBeBlank)
		})
	})
}
