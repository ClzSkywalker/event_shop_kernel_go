package test

import (
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTeam(t *testing.T) {
	ctx := initGormAndVar()
	Convey("team", t, func() {
		tmList, err := infrastructure.TeamFindMyTeam(ctx)
		So(err, ShouldBeNil)
		So(len(tmList), ShouldBeGreaterThan, 0)
		Convey("create team", func() {
			service.TeamCreate(ctx, entity.TeamCreateReq{})
		})
	})
}
