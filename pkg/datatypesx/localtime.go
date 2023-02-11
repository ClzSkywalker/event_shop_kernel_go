package datatypesx

// gorm 有一个datatypes，加上x用于区分
import (
	"database/sql/driver"
	"time"

	"github.com/clz.skywalker/event.shop/kernal/pkg/consts"
)

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+string(consts.DateTimeLayout)+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(string(consts.DateTimeLayout))+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, string(consts.DateTimeLayout))
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(string(consts.DateTimeLayout))), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(string(consts.DateTimeLayout))
}
