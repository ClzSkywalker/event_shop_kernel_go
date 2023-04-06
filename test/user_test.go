package test

import (
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserRegisterAndLogin(t *testing.T) {
	ctx1 := initGormAndVar()

	Convey("uid:register,login,bind", t, func() {
		Convey("success", func() {
			// register by uid
			remailreq := entity.RegisterByEmailReq{
				Email: email1,
			}
			remailreq.NickName = nikeName1
			remailreq.Pwd = pwd1
			token, err := service.UserRegisterByEmail(ctx1,
				remailreq)
			So(err, ShouldBeNil)
			So(token, ShouldNotBeEmpty)
			refreshDB(ctx1)

			// login by uid success
			um, err := infrastructure.LoginByUid(ctx1, entity.LoginByUidReq{Uid: ctx1.UID})
			So(err, ShouldBeNil)
			So(um.CreatedBy, ShouldEqual, ctx1.UID)

			// get user info
			um, err = infrastructure.GetUserInfo(ctx1, ctx1.UID)
			So(err, ShouldBeNil)
			So(um.CreatedBy, ShouldEqual, ctx1.UID)

			// update user
			nickName := "nk1"
			err = infrastructure.UserUpdate(ctx1, model.UserModel{CreatedBy: ctx1.UID, NickName: nickName})
			So(err, ShouldBeNil)
			refreshDB(ctx1)

			// bind email
			beReq := entity.BindEmailReq{
				Email: email1,
			}
			err = service.UserBindEmail(ctx1, beReq)
			So(err, ShouldBeNil)
			refreshDB(ctx1)

			bpEeq := entity.BindPhoneReq{
				Phone: phone1,
			}
			err = service.UserBindPhone(ctx1, bpEeq)
			So(err, ShouldBeNil)
			refreshDB(ctx1)

			um, err = infrastructure.GetUserInfo(ctx1, ctx1.UID)
			So(err, ShouldBeNil)
			So(um.NickName, ShouldEqual, nickName)
			So(um.Email, ShouldEqual, email1)
			So(um.Phone, ShouldEqual, phone1)

			// email,phone登录
			// infrastructure.LoginByEmail(ctx,entity.LoginByEmailReq{Email: email1,})

			Convey("failure", func() {
				// 绑定相同邮箱
				ctx2, err := newCtx()
				So(err, ShouldBeNil)
				err = service.UserBindEmail(ctx2, beReq)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserEmailBindByOtherErr)

				// 绑定相同电话
				err = service.UserBindPhone(ctx2, bpEeq)
				So(err.(errorx.CodeError).Code, ShouldEqual, module.UserPhoneBindByOtherErr)
			})
			// success
		})
	})
}
