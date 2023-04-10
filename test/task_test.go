package test

import (
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTask(t *testing.T) {
	ctx1 := initGormAndVar()

	Convey("task test", t, func() {
		classifyList, err := infrastructure.ClassifyFindByTeamId(ctx1)
		So(err, ShouldBeNil)
		So(len(classifyList), ShouldBeGreaterThan, 0)

		// 查询
		task, err := infrastructure.TaskFirst(ctx1, model.TaskModel{})
		So(err, ShouldBeNil)
		So(task.OnlyCode, ShouldNotBeBlank)

		// 修改
		req1 := entity.TaskUpdateReq{}
		utils.StructToStruct(task, &req1)
		req1.Title = "test1"
		err = service.TaskUpdate(ctx1, req1)
		So(err, ShouldBeNil)
		refreshDB(ctx1)

		// 查询
		taskM, err := infrastructure.TaskFirst(ctx1, model.TaskModel{OnlyCode: req1.OnlyCode})
		So(err, ShouldBeNil)
		So(taskM.Title, ShouldEqual, "test1")

		// 删除
		err = service.TaskDelete(ctx1, req1.OnlyCode)
		So(err, ShouldBeNil)
		refreshDB(ctx1)

		taskM, err = infrastructure.TaskFirst(ctx1, model.TaskModel{OnlyCode: req1.OnlyCode})
		So(err.(errorx.CodeError).Code, ShouldEqual, module.TaskNotfoundErr)
		So(taskM.OnlyCode, ShouldBeBlank)
	})
}

func TestTaskFilter(t *testing.T) {
	ctx1 := initGormAndVar()

	Convey("task filter test", t, func() {
		classifyList, err := infrastructure.ClassifyFindByTeamId(ctx1)
		So(err, ShouldBeNil)
		So(len(classifyList), ShouldBeGreaterThan, 0)

		// 全局
		// 无内容
		taskList, err := infrastructure.TaskFilter(ctx1, entity.TaskFilterParam{})
		So(err, ShouldBeNil)
		So(len(taskList), ShouldBeGreaterThan, 0)

		// 关键字
		taskList, err = infrastructure.TaskFilter(ctx1, entity.TaskFilterParam{
			Keyword: "The",
		})
		So(err, ShouldBeNil)
		So(len(taskList), ShouldBeGreaterThan, 0)

		// 时间

		// task status

		// order

		// 全部组合

		// 分类，分组
		// 无内容
		// 关键字
		// 时间

		// task status

		// order

		// 全部组合
	})
}
