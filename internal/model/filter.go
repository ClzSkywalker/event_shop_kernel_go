package model

import (
	"fmt"
	"strings"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/gorm"
)

//
// Author         : ClzSkywalker
// Date           : 2023-04-10
// Description    : task 过滤器
// return          {*}
//
type TaskFilter struct {
	db                *gorm.DB
	recursiveCondtion []ConditionTaskFilter // 父子查询调教
	condition         []ConditionTaskFilter // 通用查询条件
	order             []OrderTaskFilter
	limit             uint
	offset            uint
	sql               string
}

func NewTaskFilter(db *gorm.DB) *TaskFilter {
	return &TaskFilter{db: db}
}

//
// Author         : ClzSkywalker
// Date           : 2023-04-10
// Description    : 排序
// return          {*}
//
type OrderTaskFilter struct {
	Col       constx.TaskFilterColType
	OrderType constx.OrderType
}

//
// Author         : ClzSkywalker
// Date           : 2023-04-10
// Description    : 条件
// return          {*}
//
type ConditionTaskFilter struct {
	Operate constx.DBOpType
	Col     constx.TaskFilterColType
	ColType constx.ColType
	Value   interface{}
}

func (r *TaskFilter) AddOrder(p ...OrderTaskFilter) *TaskFilter {
	r.order = append(r.order, p...)
	return r
}

func (r *TaskFilter) AddCondition(p ...ConditionTaskFilter) *TaskFilter {
	r.condition = append(r.condition, p...)
	return r
}

func (r *TaskFilter) AddRecursiveCondition(p ...ConditionTaskFilter) *TaskFilter {
	r.recursiveCondtion = append(r.recursiveCondtion, p...)
	return r
}

func (r *TaskFilter) Limit(p uint) *TaskFilter {
	r.limit = p
	return r
}
func (r *TaskFilter) Offset(p uint) *TaskFilter {
	r.limit = p
	return r
}

func (r *TaskFilter) ToSql() (result string) {
	var innerSql strings.Builder
	innerSql.WriteString(toSqlCondition(r.condition))
	innerSql.WriteString(r.toSqlOrder())
	innerSql.WriteString(r.toSqlLF())
	recusiveSql := toSqlCondition(r.recursiveCondtion)
	r.sql = fmt.Sprintf(recursive(TaskTableName, recusiveSql)+`
	SELECT t1.*,tc.content 
	FROM all_folders t1  left join 
	%s tc on t1.content_id =tc.oc where 1=1 %s %s;`,
		TaskContentTableName, innerSql.String(), recusiveSql)
	return r.sql
}

func (r *TaskFilter) Exec() (result []entity.TaskEntity, err error) {
	r.sql = strings.ReplaceAll(r.sql, "\n", " ")
	r.sql = strings.ReplaceAll(r.sql, "\t", " ")
	err = r.db.Raw(r.sql).Scan(&result).Error
	return
}

func toSqlCondition(condition []ConditionTaskFilter) (result string) {
	if len(condition) == 0 {
		return ""
	}
	var where strings.Builder
	where.WriteString("and ")
	for i := 0; i < len(condition); i++ {
		var operate string
		var value string
		switch condition[i].Operate {
		case constx.DBOpEqual:
			operate = "="
			value = utils.ToString(condition[i].Value)
		case constx.DBOpLt:
			operate = "<"
			value = utils.ToString(condition[i].Value)
		case constx.DBOpGt:
			operate = ">"
			value = utils.ToString(condition[i].Value)
		case constx.DBOpIN:
			operate = "in"
			value = "(" + utils.ToString(condition[i].Value) + ")"
		case constx.DBOpAfterContain:
			operate = "like"
			value = utils.ToString(condition[i].Value) + "%"
		default:
			continue
		}
		if condition[i].ColType == constx.ColVarchar {
			where.WriteString(fmt.Sprintf(" %s %s '%s' ",
				condition[i].Col, operate, value))
		} else {
			where.WriteString(fmt.Sprintf(" %s %s %s ",
				condition[i].Col, operate, value))
		}
	}
	where.WriteString(" ")
	return where.String()
}

func (r *TaskFilter) toSqlOrder() string {
	if len(r.order) == 0 {
		return ""
	}
	var order strings.Builder
	order.WriteString(" order by ")
	orderList := make([]string, 0, len(r.order))
	for i := 0; i < len(r.order); i++ {
		orderList = append(orderList,
			fmt.Sprintf("%s %s", r.order[i].Col,
				r.order[i].OrderType))
	}
	order.WriteString(strings.Join(orderList, ",") + " ")
	return order.String()
}

func (r *TaskFilter) toSqlLF() string {
	var sql strings.Builder
	if r.limit > 0 {
		sql.WriteString(fmt.Sprintf(" limit %d ", r.limit))
	}
	if r.offset > 0 {
		sql.WriteString(fmt.Sprintf(" offset %d ", r.offset))
	}
	return sql.String()
}
