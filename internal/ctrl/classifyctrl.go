package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func QueryAll(c *gin.Context) {

}

func InsertClassify(c *gin.Context) {
	ret := httpx.NewResult()
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	cm := entity.ClassifyInsertReq{}
	err = validateBind(c, &cm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	value, _ := c.Get(constx.TokenUid)
	id, err := service.InsertClassify(container.GlobalServerContext.ClassifyModel, &model.ClassifyModel{
		CreatedBy: value.(string),
		Title:     cm.Title,
		Color:     cm.Color,
		Sort:      cm.Sort,
	})
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.CommonResponseId{Id: id}
}
