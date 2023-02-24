package constx

const (
	KernelVersion = "v0.0.1"

	LangEnglish = "en"
	LangChinese = "zh"

	CmdPort    = "port"
	CmdMode    = "mode"
	CmdDbPath  = "dbPath"
	CmdLogPath = "logPath"
)

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
