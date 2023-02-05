package model

type TaskModeType int

const (
	Normal TaskModeType = iota
	Day
	Week
	Month
	Year
	Workday          // 工作日(周一-周五)
	LegalWorkingDay  // 法定工作日
	StatutoryHoliday // 法定节假日
)

type TaskModeModel struct {
	ModeId      TaskModeType `json:"mode_id" gorm:"mode_id"` // 重复模式 TaskModeEnum
	ConfigBytes []byte       `json:"config" gorm:"config"`
	Config      TaskModeConfigModel
}

type TaskModeConfigModel struct {
	Day []int `json:"day" gorm:"day"`
}
