package test

import (
	"fmt"
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDevide(t *testing.T) {
	ctx1 := initGormAndVar()
	fmt.Println("testaaaaaa")
	Convey("devide success", t, func() {
		// 查
		devide, err := infrastructure.DevideFirst(ctx1, model.DevideModel{CreatedBy: ctx1.UID})
		So(err, ShouldBeNil)
		So(devide.Id, ShouldBeGreaterThan, 0)

		// 改
		title1 := "test1"
		err = service.DevideUpdate(ctx1, entity.DevideUpdateReqEntity{
			OnlyCode: devide.OnlyCode,
			Sort:     5,
			Title:    title1,
		})
		So(err, ShouldBeNil)
		refreshDB(ctx1)
		devide, err = infrastructure.DevideFirst(ctx1, model.DevideModel{OnlyCode: devide.OnlyCode})
		So(err, ShouldBeNil)
		So(devide.Sort, ShouldEqual, 5)
		So(devide.Title, ShouldEqual, title1)
		// 不变信息的改
		err = service.DevideUpdate(ctx1, entity.DevideUpdateReqEntity{
			OnlyCode: devide.OnlyCode,
			Sort:     5,
			Title:    title1,
		})
		So(err, ShouldBeNil)
		refreshDB(ctx1)

		// 增
		cm, err := infrastructure.ClassifyFirst(ctx1, model.ClassifyModel{CreatedBy: ctx1.UID})
		So(err, ShouldBeNil)
		title := "test2"
		oc, err := service.DevideInsert(ctx1, entity.DevideInsertReqEntity{
			Title:      title,
			Sort:       3,
			ClassifyId: cm.OnlyCode,
		})
		So(err, ShouldBeNil)
		So(len(oc), ShouldBeGreaterThan, 0)
		refreshDB(ctx1)
		devide, err = infrastructure.DevideFirst(ctx1, model.DevideModel{CreatedBy: ctx1.UID, Title: title})
		So(err, ShouldBeNil)
		So(devide.Title, ShouldEqual, title)
		refreshDB(ctx1)

		// 删
		err = service.DevideDelete(ctx1, oc)
		So(err, ShouldBeNil)
		refreshDB(ctx1)
	})
	Convey("devide failure", t, func() {
		ctx2, err := newCtx()
		So(err, ShouldBeNil)

		// 增 同名
		devide, err := infrastructure.DevideFirst(ctx2, model.DevideModel{CreatedBy: ctx2.UID})
		So(err, ShouldBeNil)
		So(devide.Id, ShouldBeGreaterThan, 0)
		_, err = service.DevideInsert(ctx2, entity.DevideInsertReqEntity{
			Title:      devide.Title,
			ClassifyId: devide.ClassifyId,
			Sort:       2,
		})
		So(errorx.Is(err, module.DevideTitleRepeatErr), ShouldBeTrue)
		refreshDB(ctx2)

		// 增 classifyId 不存在
		_, err = service.DevideInsert(ctx2, entity.DevideInsertReqEntity{
			Title:      "test3",
			ClassifyId: "notfound",
			Sort:       2,
		})
		So(errorx.Is(err, module.ClassifyNotfoundErr), ShouldBeTrue)
		refreshDB(ctx2)

		// 改 同名
		title2 := "repeat"
		_, err = service.DevideInsert(ctx2, entity.DevideInsertReqEntity{
			Title:      title2,
			ClassifyId: devide.ClassifyId,
			Sort:       2,
		})
		So(err, ShouldBeNil)
		refreshDB(ctx2)
		err = service.DevideUpdate(ctx2, entity.DevideUpdateReqEntity{
			OnlyCode: devide.OnlyCode,
			Title:    title2,
		})
		So(errorx.Is(err, module.DevideTitleRepeatErr), ShouldBeTrue)
		refreshDB(ctx2)

		// 删 存在task
		_, err = service.TaskInsert(ctx2, entity.TaskInsertReq{
			Title:    "task",
			DevideId: devide.OnlyCode,
		})
		So(err, ShouldBeNil)
		refreshDB(ctx2)
		err = service.DevideDelete(ctx2, devide.OnlyCode)
		So(errorx.Is(err, module.DevideDelExistTaskErr), ShouldBeTrue)
		refreshDB(ctx2)
	})
}
