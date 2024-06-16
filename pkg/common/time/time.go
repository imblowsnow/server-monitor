package time

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	if time.Time(t).IsZero() {
		return ""
	}
	return time.Time(t).Format(timeFormart)
}
func (date *Time) Scan(value interface{}) (err error) {
	// 如果为 int
	if v, ok := value.(int64); ok {
		*date = Time(time.Unix(v, 0))
		return
	}
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Time(nullTime.Time)
	return
}

func (date Time) Value() (driver.Value, error) {
	t := time.Time(date)
	if t.IsZero() {
		return nil, nil
	}
	return date.String(), nil
}

// GormDataType gorm common data type
func (date Time) GormDataType() string {
	return "datetime"
}
