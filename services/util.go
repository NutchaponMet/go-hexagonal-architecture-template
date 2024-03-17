package services

import (
	"go-hexagonal-architecture/logs"
	"time"
)

func GenStrTimeStampStartToEndOneDayPeriod(dateStr string) (start_period string, end_period string, err error) {
	date, err := time.Parse("2006-01-02", dateStr)

	if err != nil {
		logs.Error(err)
		return "", "", err
	}

	year, month, day := date.Date()

	lastHour := time.Date(year, month, day, 23, 59, 59, 0, time.Local)

	start_period = date.Format("2006-01-02T15:04:05+07:00")
	end_period = lastHour.Format("2006-01-02T15:04:05+07:00")

	return
}
