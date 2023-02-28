package constx

const (
	KernelVersion = "v0.0.1"

	LangEnglish = "en"
	LangChinese = "zh"

	PwdSalt = "LeagueOfLegends"

	CmdPort    = "port"
	CmdMode    = "mode"
	CmdDbPath  = "dbPath"
	CmdLogPath = "logPath"

	TokenIssuer    = "ClzSkywalker"
	TokenExpiresAt = 86400 * 3
	TokenSub       = "event shop"
	TokenUid       = "uid"
	TokenSecret    = "Token=5RNEYJXWQA"
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
	EmailRT    RegisterTypt = 0
	PhoneRT    RegisterTypt = 1
	WechatRT   RegisterTypt = 2
	QQRT       RegisterTypt = 3
	GoogleRT   RegisterTypt = 4
	FacebookRT RegisterTypt = 5
)
