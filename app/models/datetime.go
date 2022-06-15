package models

import (
	"database/sql/driver"
	"fmt"
	"gautam/server/app/constants"
	"math"
	"strings"
	"time"
)

//TimeLayout enum type
type TimeLayout string

//TimeLayout enum
const (
	TimeLayoutTimeSQL     TimeLayout = "15:04:05"
	TimeLayoutDateSQL     TimeLayout = "2006-01-02"
	TimeLayoutDateTimeSQL TimeLayout = "2006-01-02 15:04:05"
)

//Time format
type Time time.Time

//Date format
type Date time.Time

//DateTime format
type DateTime time.Time

func (timelayout TimeLayout) String() string {
	return string(timelayout)
}

// MarshalJSON writes a quoted string in the custom format
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

// String returns the time in the custom format
func (t *Time) String() string {
	time := time.Time(*t)
	return fmt.Sprintf("%q", time.Format(string(TimeLayoutTimeSQL)))
}

//UnmarshalJSON time format
func (t *Time) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse(string(TimeLayoutTimeSQL), strInput)
	if err != nil {
		return err
	}

	*t = Time(newTime)
	return nil
}

// Value for db column -> write
func (t Time) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(string(TimeLayoutTimeSQL))), nil
}

// Scan db time value -> read
func (t *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = Time(v)
	case nil:
		*t = Time{}
	default:
		return fmt.Errorf("cannot sql.Scan() Time from: %#v", v)
	}
	return nil
}

// UnmarshalText from db
func (t *Time) UnmarshalText(value string) error {
	dd, err := time.Parse(string(TimeLayoutTimeSQL), value)
	if err != nil {
		return err
	}
	*t = Time(dd)
	return nil
}

//------------ Date ------------ //

// MarshalJSON writes a quoted string in the custom format
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

// String returns the time in the custom format
func (d *Date) String() string {
	date := time.Time(*d)
	return fmt.Sprintf("%q", date.Format(string(TimeLayoutDateSQL)))
}

//UnmarshalJSON date format
func (d *Date) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse(string(TimeLayoutDateSQL), strInput)
	if err != nil {
		return err
	}

	*d = Date(newTime)
	return nil
}

// Value for db column -> write
func (d Date) Value() (driver.Value, error) {
	return driver.Value(time.Time(d).Format(string(TimeLayoutDateSQL))), nil
}

// Scan db time value -> read
func (d *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return d.UnmarshalText(string(v))
	case string:
		return d.UnmarshalText(v)
	case time.Time:
		*d = Date(v)
	case nil:
		*d = Date{}
	default:
		return fmt.Errorf("cannot sql.Scan() Date from: %#v", v)
	}
	return nil
}

// UnmarshalText from db
func (d *Date) UnmarshalText(value string) error {
	dd, err := time.Parse(string(TimeLayoutDateSQL), value)
	if err != nil {
		return err
	}
	*d = Date(dd)
	return nil
}

//------------ DateTime ------------ //

// MarshalJSON writes a quoted string in the custom format
func (dt DateTime) MarshalJSON() ([]byte, error) {
	return []byte(dt.String()), nil
}

// String returns the time in the custom format
func (dt *DateTime) String() string {
	datetime := time.Time(*dt)
	return fmt.Sprintf("%q", datetime.Format(string(TimeLayoutDateTimeSQL)))
}

//UnmarshalJSON date time format
func (dt *DateTime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse(string(TimeLayoutDateTimeSQL), strInput)
	if err != nil {
		return err
	}

	*dt = DateTime(newTime)
	return nil
}

// Value for db column -> write
func (dt DateTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(dt).Format(string(TimeLayoutDateTimeSQL))), nil
}

// Scan db time value -> read
func (dt *DateTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return dt.UnmarshalText(string(v))
	case string:
		return dt.UnmarshalText(v)
	case time.Time:
		*dt = DateTime(v)
	case nil:
		*dt = DateTime{}
	default:
		return fmt.Errorf("cannot sql.Scan() DateTime from: %#v", v)
	}
	return nil
}

// UnmarshalText from db
func (dt *DateTime) UnmarshalText(value string) error {
	dd, err := time.Parse(string(TimeLayoutDateTimeSQL), value)
	if err != nil {
		return err
	}
	*dt = DateTime(dd)
	return nil
}

func (dt DateTime) FormatdddMMMDo() string {
	time := time.Time(dt)

	suffix := "th"

	switch time.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}

	return time.Format("Mon, Jan 2" + suffix)
}

func (dt DateTime) FormatdddDoMMM() string {
	time := time.Time(dt)

	suffix := "th"

	switch time.Day() {
	case 1, 21, 31:
		suffix = "st"
	case 2, 22:
		suffix = "nd"
	case 3, 23:
		suffix = "rd"
	}

	return fmt.Sprintf(time.Format("Mon, 2%s Jan"), suffix)
}

func (dt DateTime) FormatDayName() string {
	return time.Time(dt).Format("Mon")
}

func (dt DateTime) FormatYYYYMMDD() string {
	return time.Time(dt).Format(string(TimeLayoutDateSQL))
}

func (dt DateTime) FormatdddMMMDD() string {
	return time.Time(dt).Format("Mon, Jan 02")
}

func (dt DateTime) In(location *time.Location) DateTime {
	if location == nil {
		return dt
	}

	return DateTime(time.Time(dt).In(location))
}

// Unix returns t as a Unix time, the number of seconds elapsed
func (dt DateTime) Unix() int64 {
	return time.Time(dt).Unix()
}

func (dt DateTime) FormatdddMMMDohmmAzz() string {
	// ddd, MMM Do, h:mm A zz
	dm := dt.FormatdddMMMDo()
	dh := time.Time(dt).Format("3:04 PM MST")
	return dm + ", " + dh
}

func (dt DateTime) IsInFuture() bool {
	return time.Time(dt).After(time.Now())
}

func (dt DateTime) IsInPast() bool {
	return time.Time(dt).Before(time.Now())
}

func (dt DateTime) FormatHHmmss() string {
	return time.Time(dt).Format("15:04:05")
}

func (dt DateTime) FormathmmA() string {
	return time.Time(dt).Format("3:04 PM")
}

func (dt DateTime) Formathmm() string {
	return time.Time(dt).Format("3:04")
}

func (dt DateTime) Formatddd() string {
	return time.Time(dt).Format("Mon")
}

func (dt DateTime) FormatYYYYMMDDTHHmmss000Z() string {
	return time.Time(dt).Format("2006-01-02T15:04:05.000Z")
}

func (dt DateTime) Format(format string) string {
	return time.Time(dt).Format(format)
}

func (dt DateTime) FormatDays(tomorrow bool, dayAfter bool) string {
	eventTime := time.Time(dt)
	timeNow := time.Now().In(eventTime.Location())

	ey, em, ed := eventTime.Date()
	ny, nm, nd := timeNow.Date()

	if ey == ny && em == nm && ed == nd {
		return constants.Today
	} else if !tomorrow {
		return ""
	}

	ty, tm, td := timeNow.Add(24 * time.Hour).Date()

	if ey == ty && em == tm && ed == td {
		return constants.Tomorrow
	} else if !dayAfter {
		return ""
	}

	ay, am, ad := timeNow.Add(48 * time.Hour).Date()

	if ey == ay && em == am && ed == ad {
		return constants.DayAfter
	}

	return ""
}

func (dt1 DateTime) UntilHours() float64 {
	t := time.Time(dt1)
	return time.Until(t).Hours()
}

// DiffInDays returns the duration td-ud. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
func (td DateTime) DiffInDays(ud DateTime) int {
	t := time.Time(td)
	u := time.Time(ud)
	return int(math.Ceil(t.Sub(u).Hours() / 24))
}

// td > ud => more than 0 else less than 0
func (td DateTime) DiffInSeconds(ud DateTime) int64 {
	t := time.Time(td)
	u := time.Time(ud)
	return int64(t.Sub(u).Seconds())
}

func (dt DateTime) IsSameTime(dt2 DateTime) bool {
	now := time.Time(dt)
	then := time.Time(dt2)

	return now.Hour() == then.Hour() && now.Minute() == then.Minute()
}

func (dt DateTime) Month() int {
	return int(time.Time(dt).Month())
}

func (dt DateTime) Weekday() int {
	return int(time.Time(dt).Weekday())
}

func (dt DateTime) Minute() int {
	return time.Time(dt).Minute()
}

func (dt DateTime) Day() int {
	return time.Time(dt).Day()
}

func (dt DateTime) FormatDMMMYYYY() string {
	t1 := time.Time(dt)
	return fmt.Sprintf("%d %s", t1.Day(), t1.Format("Jan 2006"))
}

func (dt DateTime) FormatDMMM() string {
	t1 := time.Time(dt)
	return fmt.Sprintf("%d %s", t1.Day(), t1.Format("Jan"))
}

func (dt DateTime) Hour() int {
	return time.Time(dt).Hour()
}

func (dt DateTime) AddDate(y int, m int, d int) DateTime {
	return DateTime(time.Time(dt).AddDate(y, m, d))
}

func (dt DateTime) AddMinutues(m int) DateTime {
	return DateTime(time.Time(dt).Add(time.Duration(m) * time.Minute))
}

func (dt DateTime) StartOfDay() DateTime {
	ot := time.Time(dt)
	year, month, day := ot.Date()
	nt := time.Date(year, month, day, 0, 0, 0, 0, ot.Location())
	return DateTime(nt)
}

func (dt DateTime) EndOfDay() DateTime {
	ot := time.Time(dt)
	year, month, day := ot.Date()
	nt := time.Date(year, month, day, 23, 59, 59, 0, ot.Location())
	return DateTime(nt)
}
