package constx

// 数据库初始化状态
type DbInitStateType int

type TaskModeType int

const (
	TaskModeNormal TaskModeType = iota
	TaskModeDay
	TaskModeWeek
	TaskModeMonth
	TaskModeYear
	TaskModeWorkday          // 工作日(周一-周五)
	TaskModeLegalWorkingDay  // 法定工作日
	TaskModeStatutoryHoliday // 法定节假日
	TaskModeCustome          // 自定义
)

// 任务展示模式
type TaskShowType int

const (
	TaskShowNormal TaskShowType = iota
	TaskShowSimple              // 简约模式
)

// 排序处理
type OrderType string

const (
	OrderAsc  OrderType = "asc"  // 升序
	OrderDesc OrderType = "desc" // 降序
)

// 过滤筛选
type DBOpType string

const (
	DBOpEqual            DBOpType = "equal"            // =
	DBOpNotEqual         DBOpType = "notEqual"         // !=
	DBOpGt               DBOpType = "gt"               // >
	DBOpGe               DBOpType = "ge"               // >=
	DBOpLt               DBOpType = "lt"               // <
	DBOpLe               DBOpType = "le"               // <=
	DBOpContain          DBOpType = "contain"          // 包含 Like
	DBOpNotContain       DBOpType = "notContain"       // 不包含 not like
	DBOpNull             DBOpType = "null"             // 为空
	DBOpNotNull          DBOpType = "notNull"          // 不为空
	DBOpSearch           DBOpType = "search"           // 搜索
	DBOpIN               DBOpType = "in"               // in or any
	DBOpNotIn            DBOpType = "notIn"            // not in or not any
	DBOpRange            DBOpType = "range"            // between and
	DBOpNotBelong        DBOpType = "notBelong"        // 不属于
	DBOpBeforeContain    DBOpType = "beforeContain"    // 前缀匹配 Like X%（属于contain的子集）
	DBOpAfterContain     DBOpType = "afterContain"     // 后缀匹配 Like %X（属于contain的子集
	DBOpNotBeforeContain DBOpType = "notBeforeContain" // 前缀不匹配 NOT Like X%
	DBOpNotAfterContain  DBOpType = "notAfterContain"  // 后缀不匹配 NOT Like %X
)

// 列类型
type ColType string

const (
	ColVarchar ColType = "string"
	ColInteger ColType = "INTEGER"
)

type TaskFilterColType string

const (
	TaskFilterColCreatedAt   TaskFilterColType = "t1.created_at"
	TaskFilterColDevideId    TaskFilterColType = "t1.devide_id"
	TaskFilterColCreatedBy   TaskFilterColType = "t1.created_by"
	TaskFilterColTitle       TaskFilterColType = "t1.title"
	TaskFilterColEndAt       TaskFilterColType = "t1.end_at"
	TaskFilterColCompletedAt TaskFilterColType = "t1.completed_at"
)

// 任务状态
type TaskCompleteType int

const (
	TaskCompleteAll      TaskCompleteType = iota // 进行中，完成
	TaskCompleteUnderWay                         // 进行中
	TaskCompleted                                // 完成
)

// 任务排序模式
type TaskOrderType int

const (
	TaskOrderDefault  = iota // 默认，创建时间排序
	TaskOrderGroup           // 分组排序,创建时间
	TaskOrderEndTime         // 截止时间排序
	TaskOrderImpotant        // 重要程度排序，创建时间
)

// 时间格式
type DateCompType string

const (
	DateTimeLayoutTimeZone DateCompType = "2006-01-02 15:04:05+08"
	DateTimeLayout         DateCompType = "2006-01-02 15:04:05"
	DateTime2Layout        DateCompType = "2006-01-02 15:04" // 年-月-日 时：分
	DateYMD                DateCompType = "2006-01-02"       // 年-月-日
	DateYM                 DateCompType = "2006-01"          // 年-月
	DateMD                 DateCompType = "01-02"            // 月-日
	DateY                  DateCompType = "2006"             // 年
	DateQ                  DateCompType = "quarter"          // 季度
	DateM                  DateCompType = "01"               // 月
	DateD                  DateCompType = "02"               // 天
)

// 用户类型
type UserType int

const (
	NormalUT    UserType = 0 // 普通成员
	MemberUT    UserType = 1 // 有时限的会员
	PermanentUT UserType = 2 // 永久会员
)

// 用户注册方式
type RegisterTypt int

const (
	EmailRT RegisterTypt = 1
	PhoneRT RegisterTypt = 2
	// WechatRT   RegisterTypt = 2
	// QQRT       RegisterTypt = 3
	// GoogleRT   RegisterTypt = 4
	// FacebookRT RegisterTypt = 5
)
