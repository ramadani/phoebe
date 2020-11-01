package phoebe

import (
	"time"
)

type Carbon interface {
	Now() time.Time
	StartOfWeek() time.Time
	EndOfWeek() time.Time
	StartOfMonth() time.Time
	EndOfMonth() time.Time

	StartOfWeekday() time.Time
	EndOfWeekday() time.Time
	StartOfWeekend() time.Time
	EndOfWeekend() time.Time
	IsWeekday() bool
	IsWeekend() bool

	Of(t time.Time) time.Time
	StartOfDayOf(t time.Time) time.Time
	EndOfDayOf(t time.Time) time.Time

	StartOfWeekOf(t time.Time) time.Time
	EndOfWeekOf(t time.Time) time.Time
	StartOfMonthOf(t time.Time) time.Time
	EndOfMonthOf(t time.Time) time.Time

	StartOfWeekdayOf(t time.Time) time.Time
	EndOfWeekdayOf(t time.Time) time.Time
	StartOfWeekendOf(t time.Time) time.Time
	EndOfWeekendOf(t time.Time) time.Time
	IsWeekdayOf(t time.Time) bool
	IsWeekendOf(t time.Time) bool
}
