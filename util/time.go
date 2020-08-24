package util

import "time"

// StringToTime is a basic converter from a string date to a time.Time.
func StringToTime(value string) time.Time {
	layout := "2006-01-02T15:04:05"
	convertedTime, _ := time.Parse(layout, value)

	return convertedTime
}
