package test

import (
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserRegister(t *testing.T) {
	initGormAndVar()
	Convey("user login and register", t, func() {
		Convey("register by emial", func() {
			remailreq := entity.RegisterByEmailReq{
				Email: email1,
			}
			remailreq.NickName = nikeName1
			remailreq.Pwd = pwd1
			uid, err := service.RegisterByEmail(container.GlobalServerContext.UserModel, remailreq)
			So(err, ShouldBeNil)
			So(uid, ShouldNotBeEmpty)
			_, err = service.RegisterByEmail(container.GlobalServerContext.UserModel, remailreq)
			So(err, ShouldNotBeNil)
		})

		Convey("login by email", func() {
			req := entity.LoginByEmailReq{
				Email: email1,
				Pwd:   pwd1,
			}
			Convey("success", func() {
				uid, err := service.LoginByEmail(container.GlobalServerContext.UserModel, req)
				So(err, ShouldBeNil)
				So(uid, ShouldNotBeEmpty)
			})
			Convey("pwd error", func() {
				req.Pwd += "error"
				uid, err := service.LoginByEmail(container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserPwdErr)
				So(uid, ShouldBeBlank)
			})
			Convey("notfound", func() {
				req.Email = "notfound@qq.com"
				uid, err := service.LoginByEmail(container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserNotFoundErr)
				So(uid, ShouldBeEmpty)
			})
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

		Convey("login by phone", func() {
			req := entity.LoginByPhoneReq{Phone: phone1}
			req.Pwd = pwd1
			Convey("success", func() {
				uid, err := service.LoginByPhone(container.GlobalServerContext.UserModel, req)
				So(err, ShouldBeNil)
				So(uid, ShouldNotBeEmpty)
			})
			Convey("pwd error", func() {
				req.Pwd += "error"
				uid, err := service.LoginByPhone(container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserPwdErr)
				So(uid, ShouldBeBlank)
			})
			Convey("notfound", func() {
				req.Phone = "notfound"
				uid, err := service.LoginByPhone(container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserNotFoundErr)
				So(uid, ShouldBeEmpty)
			})
		})

		Convey("register by uid", func() {
			uid, err := service.RegisterByUid(container.GlobalServerContext.UserModel)
			So(err, ShouldBeNil)
			So(uid, ShouldNotBeBlank)
		})
	})
}

func TestUserLogin(t *testing.T) {
	initGormAndVar()
	Convey("user login", t, func() {
		Convey("login by emial", func() {

		})
	})
}
