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
