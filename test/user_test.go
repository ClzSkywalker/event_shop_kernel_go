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
	ctx := initGormAndVar()
	Convey("user login, register, bind", t, func() {
		Convey("register by emial", func() {
			remailreq := entity.RegisterByEmailReq{
				Email: email1,
			}
			remailreq.NickName = nikeName1
			remailreq.Pwd = pwd1
			um, err := service.RegisterByEmail(ctx, container.GlobalServerContext.UserModel,
				container.GlobalServerContext.TeamModel, remailreq)
			So(err, ShouldBeNil)
			So(um.CreatedBy, ShouldNotBeEmpty)
			So(um.TeamIdPort, ShouldNotBeEmpty)
			_, err = service.RegisterByEmail(ctx, container.GlobalServerContext.UserModel,
				container.GlobalServerContext.TeamModel, remailreq)
			So(err, ShouldNotBeNil)
		})

		Convey("login by email", func() {
			req := entity.LoginByEmailReq{
				Email: email1,
				Pwd:   pwd1,
			}
			Convey("success", func() {
				uid, err := service.LoginByEmail(ctx, container.GlobalServerContext.UserModel, req)
				So(err, ShouldBeNil)
				So(uid, ShouldNotBeEmpty)
			})
			Convey("pwd error", func() {
				req.Pwd += "error"
				_, err := service.LoginByEmail(ctx, container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserPwdErr)
			})
			Convey("notfound", func() {
				req.Email = "notfound@qq.com"
				_, err := service.LoginByEmail(ctx, container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserNotFoundErr)
			})
		})

		Convey("register by phone", func() {
			req := entity.RegisterByPhoneReq{Phone: phone1}
			req.NickName = nikeName1
			req.Pwd = pwd1
			um, err := service.RegisterByPhone(ctx, container.GlobalServerContext.UserModel,
				container.GlobalServerContext.TeamModel, req)
			So(err, ShouldBeNil)
			So(um.CreatedBy, ShouldNotBeEmpty)
			So(um.TeamIdPort, ShouldNotBeEmpty)
			_, err = service.RegisterByPhone(ctx, container.GlobalServerContext.UserModel,
				container.GlobalServerContext.TeamModel, req)
			So(err, ShouldNotBeNil)
		})

		Convey("login by phone", func() {
			req := entity.LoginByPhoneReq{Phone: phone1}
			req.Pwd = pwd1
			Convey("success", func() {
				um, err := service.LoginByPhone(ctx, container.GlobalServerContext.UserModel, req)
				So(err, ShouldBeNil)
				So(um.CreatedBy, ShouldNotBeEmpty)
				So(um.TeamIdPort, ShouldNotBeEmpty)
			})
			Convey("pwd error", func() {
				req.Pwd += "error"
				_, err := service.LoginByPhone(ctx, container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserPwdErr)
			})
			Convey("notfound", func() {
				req.Phone = "notfound"
				_, err := service.LoginByPhone(ctx, container.GlobalServerContext.UserModel, req)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserNotFoundErr)
			})
		})

		Convey("register by uid", func() {
			um, err := service.RegisterByUid(ctx, container.GlobalServerContext.UserModel,
				container.GlobalServerContext.TeamModel)
			So(err, ShouldBeNil)
			So(um.CreatedBy, ShouldNotBeEmpty)
			So(um.TeamIdPort, ShouldNotBeEmpty)
			Convey("login by uid", func() {
				luq := entity.LoginByUidReq{Uid: um.CreatedBy}
				Convey("success", func() {
					um, err = service.LoginByUid(ctx, container.GlobalServerContext.UserModel, luq)
					So(err, ShouldBeNil)
					So(um.CreatedAt, ShouldNotBeEmpty)
				})
				Convey("notfound", func() {
					luq.Uid = "notfound"
					_, err = service.LoginByUid(ctx, container.GlobalServerContext.UserModel, luq)
					So(err.(errorx.CodeError).Code, ShouldEqual, module.UserNotFoundErr)
				})
			})
		})

		Convey("bind user", func() {
			um, err := service.RegisterByUid(ctx, container.GlobalServerContext.UserModel,
				container.GlobalServerContext.TeamModel)
			So(err, ShouldBeNil)
			So(um.CreatedBy, ShouldNotBeEmpty)
			So(um.TeamIdPort, ShouldNotBeEmpty)
			Convey("bind by email", func() {
				req := entity.BindEmailReq{Email: email2}
				Convey("success", func() {
					err = service.BindEmailByUid(ctx, container.GlobalServerContext.UserModel, um.CreatedBy, req)
					So(err, ShouldBeNil)
				})

				Convey("uid not found", func() {
					err = service.BindEmailByUid(ctx, container.GlobalServerContext.UserModel, "notfound", req)
					So(err.(errorx.CodeError).Code, ShouldEqual, module.UserNotFoundErr)
				})

				Convey("uid binded email", func() {
					req = entity.BindEmailReq{Email: email3}
					err = service.BindEmailByUid(ctx, container.GlobalServerContext.UserModel, um.CreatedBy, req)
					So(err, ShouldBeNil)
					req.Email = "notfound"
					err = service.BindEmailByUid(ctx, container.GlobalServerContext.UserModel, um.CreatedBy, req)
					So(err.(errorx.CodeError).Code, ShouldEqual, module.UserBindedEmailErr)
				})

				Convey("The mailbox is in use", func() {
					err = service.BindEmailByUid(ctx, container.GlobalServerContext.UserModel, um.CreatedBy, req)
					So(err.(errorx.CodeError).Code, ShouldEqual, module.UserEmailBindByOtherErr)
				})
			})
		})
	})
}
