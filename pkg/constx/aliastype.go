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
	TaskShowSimple = iota + 1 // 简约模式
)

// 任务排序模式
type TaskOrderType int

const (
	TaskOrderDefault  = iota // 默认，创建时间排序
	TaskOrderGroup           // 分组排序,创建时间
	TaskOrderEndTime         // 截止时间排序
	TaskOrderImpotant        // 重要程度排序，创建时间
)
