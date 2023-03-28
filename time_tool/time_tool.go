package time_tool

import "time"

// 获取至次日凌晨零点与当前时间的时间差
func Get2NextDayZeroTimeSecond() time.Duration {
	theNextDay := time.Now().AddDate(0, 0, 1)
	theNDZeroTime := time.Date(theNextDay.Year(), theNextDay.Month(), theNextDay.Day(), 0, 0, 0, 0, theNextDay.Location())
	//return theNDZeroTime.Sub(now).Seconds()
	return time.Until(theNDZeroTime)
}
