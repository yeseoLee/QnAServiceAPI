package util

import (
	"log"
	"time"
)

const (
	formatTime     = "15:04:05"
	formatDate     = "2006-01-02"
	formatDateTime = "2006-01-02 15:04:05"
)

const (
	LOCATION_SEOUL = "Asia/Seoul"
)

var location *time.Location

func init() {
	myloc, err := time.LoadLocation(LOCATION_SEOUL)
	if err != nil {
		log.Fatal(err)
	}
	location = myloc
}

func DateTimeNow() string {
	now := time.Now().In(location)
	return now.Format(formatDateTime)
}

func DateTimeStringToTime(dateTime string) time.Time {
	t, err := time.Parse(formatDateTime, dateTime)
	if err != nil {
		log.Print(err)
	}
	return t
}
func TimeToDateTimeString(t time.Time) (dateTime string) {
	return t.Format(formatDateTime)
}
