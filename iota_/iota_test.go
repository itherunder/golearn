package iota_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

type Weekday int

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (weekDay Weekday) ToString() string {
	switch weekDay {
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	case Sunday:
		return "Sunday"
	default:
		return Weekday(int(weekDay) % 7).ToString()
	}
}

func TestIota(t *testing.T) {
	for i := 0; i < 7; i++ {
		colorlog.Info("Today is %v\n", Weekday(i).ToString())
	}
}
