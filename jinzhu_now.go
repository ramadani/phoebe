package phoebe

import (
	"github.com/jinzhu/now"
	"time"
)

type jinzhuNow struct {
	config         *now.Config
	offsetDuration time.Duration
}

func (n jinzhuNow) StartOfWeekday() time.Time {
	return n.StartOfWeekdayOf(n.Now())
}

func (n jinzhuNow) StartOfWeekdayOf(tm time.Time) time.Time {
	t := n.config.With(tm).BeginningOfWeek()

	return n.getTime(t)
}

func (n jinzhuNow) EndOfWeekday() time.Time {
	return n.EndOfWeekdayOf(n.Now())
}

func (n jinzhuNow) EndOfWeekdayOf(tm time.Time) time.Time {
	t := now.New(n.StartOfWeekdayOf(tm).AddDate(0, 0, 4)).EndOfDay()

	return n.endOfDay(n.getTime(t))
}

func (n jinzhuNow) StartOfWeekend() time.Time {
	return n.StartOfWeekendOf(n.Now())
}

func (n jinzhuNow) StartOfWeekendOf(tm time.Time) time.Time {
	t := now.With(tm).Sunday().AddDate(0, 0, -1)

	return n.getTime(t)
}

func (n jinzhuNow) EndOfWeekend() time.Time {
	return n.EndOfWeekendOf(n.Now())
}

func (n jinzhuNow) EndOfWeekendOf(tm time.Time) time.Time {
	t := now.With(tm).EndOfSunday()

	return n.endOfDay(n.getTime(t))
}

func (n jinzhuNow) IsWeekday() bool {
	return n.IsWeekdayOf(n.Now())
}

func (n jinzhuNow) IsWeekdayOf(tm time.Time) bool {
	weekday := int(tm.Weekday())
	return weekday >= 1 && weekday <= 5
}

func (n jinzhuNow) IsWeekend() bool {
	return n.IsWeekendOf(n.Now())
}

func (n jinzhuNow) IsWeekendOf(tm time.Time) bool {
	weekday := int(tm.Weekday())
	return weekday == 6 || weekday == 0
}

func (n jinzhuNow) Now() time.Time {
	t := time.Now()

	return n.getTime(t)
}

func (n jinzhuNow) Of(tm time.Time) time.Time {
	return n.getTime(tm)
}

func (n jinzhuNow) StartOfWeek() time.Time {
	return n.StartOfWeekOf(n.Now())
}

func (n jinzhuNow) StartOfWeekOf(tm time.Time) time.Time {
	t := n.config.With(tm).BeginningOfWeek()

	return n.getTime(t)
}

func (n jinzhuNow) EndOfWeek() time.Time {
	return n.EndOfWeekOf(n.Now())
}

func (n jinzhuNow) EndOfWeekOf(tm time.Time) time.Time {
	t := n.config.With(tm).EndOfWeek()

	return n.endOfDay(n.getTime(t))
}

func (n jinzhuNow) StartOfMonth() time.Time {
	return n.StartOfMonthOf(n.Now())
}

func (n jinzhuNow) StartOfMonthOf(tm time.Time) time.Time {
	t := now.With(tm).BeginningOfMonth()

	return n.getTime(t)
}

func (n jinzhuNow) EndOfMonth() time.Time {
	return n.EndOfMonthOf(n.Now())
}

func (n jinzhuNow) EndOfMonthOf(tm time.Time) time.Time {
	t := now.With(tm).EndOfMonth()

	return n.endOfDay(n.getTime(t))
}

func (n jinzhuNow) StartOfDayOf(tm time.Time) time.Time {
	t := now.With(tm).BeginningOfDay()

	return n.getTime(t)
}

func (n jinzhuNow) EndOfDayOf(tm time.Time) time.Time {
	t := now.With(tm).EndOfDay()

	return n.endOfDay(n.getTime(t))
}

func (n jinzhuNow) endOfDay(tm time.Time) time.Time {
	return tm.Add(-1 * time.Second)
}

func (n jinzhuNow) getTime(t time.Time) time.Time {
	return t.Add(n.offsetDuration)
}

func NewJinzhuNow(weekStartDay time.Weekday, offsetHour int) Carbon {
	config := &now.Config{
		WeekStartDay: weekStartDay,
	}

	offsetDur := time.Duration(offsetHour) * time.Hour

	return &jinzhuNow{config: config, offsetDuration: offsetDur}
}
