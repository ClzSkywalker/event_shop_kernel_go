package test

import (
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClassifySuccess(t *testing.T) {
	ctx1 := initGormAndVar()
	Convey("classify success", t, func() {
		// 查
		// 查批量
		cList, err := infrastructure.ClassifyFindByTeamId(ctx1)
		So(err, ShouldBeNil)
		So(len(cList), ShouldBeGreaterThan, 0)

		// 查单个
		classify, err := infrastructure.ClassifyFirst(ctx1, model.ClassifyModel{CreatedBy: ctx1.UID})
		So(err, ShouldBeNil)
		So(classify.OnlyCode, ShouldNotBeBlank)

		// 增
		// 增父类
		oc, err := service.ClassifyCreate(ctx1, classify1)
		So(err, ShouldBeNil)
		So(oc, ShouldNotBeBlank)
		refreshDB(ctx1)
		classify, err = infrastructure.ClassifyFirst(ctx1, model.ClassifyModel{
			OnlyCode: oc,
		})
		So(err, ShouldBeNil)
		So(classify.OnlyCode, ShouldEqual, oc)

		// 增子类
		classify2c := classify2
		classify2c.ParentId = oc
		oc1, err := service.ClassifyCreate(ctx1, classify2c)
		So(err, ShouldBeNil)
		So(oc1, ShouldNotBeBlank)
		refreshDB(ctx1)

		// update
		err = service.ClassifyUpdate(ctx1, entity.ClassifyUpdateReq{
			OnlyCode:  oc,
			Title:     "test2",
			ShowType:  constx.TaskShowNormal,
			OrderType: constx.TaskOrderDefault,
			Color:     "#fffadc",
			Sort:      6,
		})
		So(err, ShouldBeNil)
		refreshDB(ctx1)
		classify, err = infrastructure.ClassifyFirst(ctx1, model.ClassifyModel{CreatedBy: ctx1.UID, OnlyCode: oc})
		So(err, ShouldBeNil)
		So(classify.OnlyCode, ShouldEqual, oc)
		So(classify.Title, ShouldEqual, "test2")

		// delete
		err = service.ClassifyDel(ctx1, oc)
		So(err, ShouldBeNil)
		refreshDB(ctx1)
		_, err = infrastructure.ClassifyFirst(ctx1, model.ClassifyModel{CreatedBy: ctx1.UID, OnlyCode: oc})
		So(errorx.Is(err, module.ClassifyNotfoundErr), ShouldBeTrue)
	})
}

func TestClasifyFailure(t *testing.T) {
	ctx1 := initGormAndVar()
	Convey("classify failure", t, func() {
		// !增
		// 同名
		oc, err := service.ClassifyCreate(ctx1, classify1)
		So(err, ShouldBeNil)
		So(oc, ShouldNotBeBlank)
		refreshDB(ctx1)
		oc, err = service.ClassifyCreate(ctx1, classify1)
		So(errorx.Is(err, module.ClassifyTitleRepeatErr), ShouldBeTrue)
		So(oc, ShouldBeBlank)
		refreshDB(ctx1)

		// teamId不存在
		tid := ctx1.TID
		ctx1.TID = "notfound"
		_, err = service.ClassifyCreate(ctx1, classify2)
		So(errorx.Is(err, module.TeamNotFound), ShouldBeTrue)
		ctx1.TID = tid
		refreshDB(ctx1)

		// 父分类不存在
		classify2c := classify2
		classify2c.ParentId = "notfound"
		_, err = service.ClassifyCreate(ctx1, classify2c)
		So(errorx.Is(err, module.ClassifyParentNoExistErr), ShouldBeTrue)
		refreshDB(ctx1)

		// 层级超过两层
		classify2c.ParentId = ""
		oc, err = service.ClassifyCreate(ctx1, classify2c)
		So(err, ShouldBeNil)
		refreshDB(ctx1)

		classify2c.Title = "title333"
		classify2c.ParentId = oc
		oc, err = service.ClassifyCreate(ctx1, classify2c)
		So(err, ShouldBeNil)
		refreshDB(ctx1)

		classify2c.Title = "title444"
		classify2c.ParentId = oc
		oc, err = service.ClassifyCreate(ctx1, classify2c)
		So(errorx.Is(err, module.ClassifyDeepErr), ShouldBeTrue)
		So(oc, ShouldBeBlank)
		refreshDB(ctx1)

		// !update
		// 同名
		ctx1, err = newCtx()
		So(err, ShouldBeNil)
		cList, err := infrastructure.ClassifyFindByTeamId(ctx1)
		So(err, ShouldBeNil)
		So(len(cList), ShouldBeGreaterThan, 0)

		err = service.ClassifyUpdate(ctx1, entity.ClassifyUpdateReq{
			OnlyCode: cList[0].OnlyCode,
			Title:    cList[2].Title,
		})
		So(errorx.Is(err, module.ClassifyTitleRepeatErr), ShouldBeTrue)
		refreshDB(ctx1)

		// team notfound
		tid = ctx1.TID
		ctx1.TID = "notfound"
		err = service.ClassifyUpdate(ctx1, entity.ClassifyUpdateReq{
			OnlyCode: cList[0].OnlyCode,
			Title:    "test1",
		})
		So(errorx.Is(err, module.TeamNotFound), ShouldBeTrue)
		ctx1.TID = tid
		refreshDB(ctx1)

		// !delete
		// 删除分类还存在devide
		err = service.ClassifyDel(ctx1, cList[0].OnlyCode)
		So(errorx.Is(err, module.ClassifyDelExistDevideErr), ShouldBeTrue)
		refreshDB(ctx1)

		// 删除不存在/已被删除
		err = service.ClassifyDel(ctx1, "notfound")
		So(err, ShouldBeNil)
	})
}
