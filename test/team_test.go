package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTeam(t *testing.T) {
	Convey("team", t, func() {

	})

	// ctx := initGormAndVar()
	// Convey("team", t, func() {
	// 	tmList, err := infrastructure.TeamFindMyTeam(ctx)
	// 	So(err, ShouldBeNil)
	// 	So(len(tmList), ShouldBeGreaterThan, 0)
	// 	Convey("create team", func() {
	// 		tid, err := service.TeamCreate(ctx, entity.TeamCreateReq{
	// 			Name:        "name",
	// 			Description: "description",
	// 			Sort:        1,
	// 		})
	// 		So(err, ShouldBeNil)
	// 		So(tid, ShouldNotBeEmpty)

	// 		ctx.BaseTx = *container.NewBaseServiceContext(container.GlobalServerContext, container.GlobalServerContext.Db)
	// 		tm, err := infrastructure.TeamQueryFirst(ctx, model.TeamModel{TeamId: tid})
	// 		So(err, ShouldBeNil)
	// 		So(tm.Name, ShouldEqual, "name")
	// 		So(tm.Description, ShouldEqual, "description")

	// 		ctx.BaseTx = *container.NewBaseServiceContext(container.GlobalServerContext, container.GlobalServerContext.Db)
	// 		err = service.TeamUpdate(ctx, entity.TeamUpdateReq{
	// 			TeamItem: entity.TeamItem{Name: "name1", Description: "d2", TeamId: tid},
	// 		})
	// 		So(err, ShouldBeNil)

	// 		ctx.BaseTx = *container.NewBaseServiceContext(container.GlobalServerContext, container.GlobalServerContext.Db)
	// 		tm, err = infrastructure.TeamQueryFirst(ctx, model.TeamModel{TeamId: tid})
	// 		So(err, ShouldBeNil)
	// 		So(tm.Name, ShouldEqual, "name1")
	// 		So(tm.Description, ShouldEqual, "d2")

	// 		tmList, err = infrastructure.TeamFindMyTeam(ctx)
	// 		So(err, ShouldBeNil)
	// 		So(len(tmList), ShouldEqual, 2)
	// 		// So(tmList[1].Sort, ShouldEqual, 1)
	// 	})
	// })
}
