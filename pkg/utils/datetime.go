package utils

import (
	"go-hexagonal/pkg/logs"
	"time"
)

// layout like "2006-01-02T15:04:05Z" or "2006-01-02 15:04:05" or "2006-01-02"
func ParseDatetime(layout, dateString string) (string, error) {
	t, err := time.Parse(layout, dateString)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	newDateString := t.Format(layout)
	return newDateString, nil
}
