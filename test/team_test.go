package test

import (
	"testing"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTeam(t *testing.T) {
	ctx := initGormAndVar()
	Convey("team", t, func() {
		tmList, err := service.TeamFindMyTeam(ctx, container.GlobalServerContext.TeamModel)
		So(err, ShouldBeNil)
		So(len(tmList), ShouldBeGreaterThan, 0)
	})
}
