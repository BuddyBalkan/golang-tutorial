package time_tool

import (
	"testing"
)

func TestGetNextDayZeroTimeStampSecond(t *testing.T) {
	toNextDateTimeZero := Get2NextDayZeroTimeSecond()
	t.Logf("the time to next day zero time is : %f \n", toNextDateTimeZero.Seconds())
	t.Logf("the time to next day zero time is : %q \n", toNextDateTimeZero)
}
