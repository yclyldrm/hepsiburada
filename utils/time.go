package utils

import (
	"time"
)

const (
	HOUR_TIME_FORMAT  = "15:04"
	START_TIME_FORMAT = "00:00"
)

var systemTime = START_TIME_FORMAT

func GetSystemTime() string {
	return systemTime
}

func IncreaseTime(add int) error {
	sysTime, err := time.Parse(HOUR_TIME_FORMAT, GetSystemTime())
	if err != nil {
		return err
	}

	sysTime = sysTime.Add(time.Hour * time.Duration(add))

	systemTime = sysTime.Format(HOUR_TIME_FORMAT)

	return nil
}

func TimeDifference(endTimeStr string) int {
	sysTime, err := time.Parse(HOUR_TIME_FORMAT, GetSystemTime())
	if err != nil {
		return 0
	}

	endTime, err := time.Parse(HOUR_TIME_FORMAT, endTimeStr)
	if err != nil {
		return 0
	}

	return int(endTime.Sub(sysTime).Hours())
}

func CalculateEndTime(duration int) string {
	sysTime, err := time.Parse(HOUR_TIME_FORMAT, GetSystemTime())
	if err != nil {
		return ""
	}

	endTime := sysTime.Add(time.Hour * time.Duration(duration))

	return endTime.Format(HOUR_TIME_FORMAT)
}

func SetStartTime() {
	systemTime = START_TIME_FORMAT
}
