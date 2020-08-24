package util

import "testing"

const errorMessage = "%s should be %d, but is %d."

func TestStringToTime(t *testing.T) {
	convertedTime := StringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf(errorMessage, "Year", 2019, convertedTime.Year())
	}

	if convertedTime.Month() != 02 {
		t.Errorf(errorMessage, "Month", 02, convertedTime.Month())
	}

	if convertedTime.Day() != 12 {
		t.Errorf(errorMessage, "Day", 12, convertedTime.Day())
	}

	if convertedTime.Hour() != 10 {
		t.Errorf(errorMessage, "Hour", 10, convertedTime.Hour())
	}

	if convertedTime.Minute() != 10 {
		t.Errorf(errorMessage, "Minute", 10, convertedTime.Minute())
	}

	if convertedTime.Second() != 10 {
		t.Errorf(errorMessage, "Second", 10, convertedTime.Second())
	}
}
