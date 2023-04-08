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

	Convey("tesk test", t, func() {
		classifyList, err := infrastructure.ClassifyFindByTeamId(ctx1)
		So(err, ShouldBeNil)
		So(len(classifyList), ShouldBeGreaterThan, 0)

		classItem := classifyList[0]
		taskList, err := infrastructure.TaskFindByClassifyId(ctx1, classItem.OnlyCode)
		So(err, ShouldBeNil)
		So(len(taskList), ShouldBeGreaterThan, 0)

		req1 := entity.TaskUpdateReq{}
		utils.StructToStruct(taskList[0], &req1)
		req1.Title = "test1"
		err = service.TaskUpdate(ctx1, req1)
		So(err, ShouldBeNil)
		refreshDB(ctx1)

		taskM, err := infrastructure.TaskFirst(ctx1, model.TaskModel{OnlyCode: req1.OnlyCode})
		So(err, ShouldBeNil)
		So(taskM.Title, ShouldEqual, "test1")

		err = service.TaskDelete(ctx1, req1.OnlyCode)
		So(err, ShouldBeNil)
		refreshDB(ctx1)

		taskM, err = infrastructure.TaskFirst(ctx1, model.TaskModel{OnlyCode: req1.OnlyCode})
		So(err.(errorx.CodeError).Code, ShouldEqual, module.TaskNotfoundErr)
		So(taskM.OnlyCode, ShouldBeBlank)
	})
}
