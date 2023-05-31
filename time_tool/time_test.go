package time_tool

import (
	"testing"
	"time"
)

func TestGetNextDayZeroTimeStampSecond(t *testing.T) {
	toNextDateTimeZero := Get2NextDayZeroTimeSecond()
	t.Logf("the time to next day zero time is : %f \n", toNextDateTimeZero.Seconds())
	t.Logf("the time to next day zero time is : %q \n", toNextDateTimeZero)
}

// 计算指定的两天之差 单位：天
func TestTimeSub2TargetDay(t *testing.T) {
	today := time.Now()
	targetDay := time.Date(today.Year(), time.February, 12, 0, 0, 0, 0, today.Location())
	result := today.Sub(targetDay)
	t.Logf("the result day is : %v", int(result.Hours()/24))
}
