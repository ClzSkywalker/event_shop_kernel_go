package infrastructure

import (
	"fmt"
	"strings"

	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
)

//
// Author         : ClzSkywalker
// Date           : 2023-04-10
// Description    : 任务过滤器
// param           {*contextx.Contextx} ctx
// param           {entity.TaskFilterParam} param
// return          {*}
//
func TaskFilter(ctx *contextx.Contextx, param entity.TaskFilterParam) (result []entity.TaskEntity,
	err error) {
	filter := model.NewTaskFilter(ctx.BaseTx.Db)

	taskSetCondition(ctx, filter, param)
	taskSetOrder(filter, param)

	filter.ToSql()
	result, err = filter.Exec()
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("param", param))
		return
	}
	return
}

func taskSetCondition(ctx *contextx.Contextx, filter *model.TaskFilter, param entity.TaskFilterParam) {
	recursiveCondtion := make([]model.ConditionTaskFilter, 1)
	condition := make([]model.ConditionTaskFilter, 0)

	// 装载 uid
	recursiveCondtion[0] = model.ConditionTaskFilter{
		Col:     constx.TaskFilterColCreatedBy,
		Operate: constx.DBOpEqual,
		ColType: constx.ColVarchar,
		Value:   ctx.UID,
	}

	// 过滤 devideIds
	if len(param.DevideIds) > 0 {
		recursiveCondtion = append(recursiveCondtion, model.ConditionTaskFilter{
			Col:     constx.TaskFilterColDevideId,
			Operate: constx.DBOpIN,
			ColType: constx.ColVarchar,
			Value:   param.DevideIds,
		})
	}

	// 根据 classifyId 获取 devideIds
	if len(param.ClassifyId) > 0 {
		condition = append(condition, model.ConditionTaskFilter{
			Col:     constx.TaskFilterColDevideId,
			Operate: constx.DBOpIN,
			ColType: constx.ColVarchar,
			Value: fmt.Sprintf(
				"select d.oc from devide d left join classify c on d.classify_id = c.oc where c.oc in (%s)",
				strings.Join(param.ClassifyId, ",")),
		})
	}

	// 关键词过滤
	if len(param.Keyword) > 0 {
		condition = append(condition, model.ConditionTaskFilter{
			Col:     constx.TaskFilterColTitle,
			Operate: constx.DBOpAfterContain,
			ColType: constx.ColVarchar,
			Value:   param.Keyword,
		})
	}

	// 范围区间筛选
	if param.BeginAt > 0 {
		recursiveCondtion = append(recursiveCondtion, model.ConditionTaskFilter{
			Col:     constx.TaskFilterColCreatedAt,
			Operate: constx.DBOpGt,
			ColType: constx.ColInteger,
			Value:   0,
		})
	}
	if param.CloseAt > 0 {
		recursiveCondtion = append(recursiveCondtion, model.ConditionTaskFilter{
			Col:     constx.TaskFilterColCreatedAt,
			Operate: constx.DBOpLt,
			ColType: constx.ColInteger,
			Value:   0,
		})
	}

	switch param.TaskStatus {
	case constx.TaskCompleteAll:
	case constx.TaskCompleteUnderWay:
		recursiveCondtion = append(recursiveCondtion, model.ConditionTaskFilter{
			Col:     constx.TaskFilterColCompletedAt,
			Operate: constx.DBOpEqual,
			ColType: constx.ColInteger,
			Value:   0,
		})
	case constx.TaskCompleted:
		recursiveCondtion = append(recursiveCondtion, model.ConditionTaskFilter{
			Col:     constx.TaskFilterColCompletedAt,
			Operate: constx.DBOpGt,
			ColType: constx.ColInteger,
			Value:   0,
		})
	}

	filter.AddCondition(condition...).AddRecursiveCondition(recursiveCondtion...)
}

func taskSetOrder(filter *model.TaskFilter, param entity.TaskFilterParam) {
	order := model.OrderTaskFilter{
		OrderType: constx.OrderDesc,
	}
	// 排序
	switch param.OrderType {
	case constx.TaskOrderDefault:
		order.Col = constx.TaskFilterColCreatedBy
	case constx.TaskOrderEndTime:
		order.Col = constx.TaskFilterColEndAt
	case constx.TaskOrderGroup, constx.TaskOrderImpotant: // 待做
		order.Col = constx.TaskFilterColCreatedBy
	default:
		order.Col = constx.TaskFilterColCreatedBy
	}
}
