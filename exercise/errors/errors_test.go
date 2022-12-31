package timeparse

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	timeString := "14:07:33"
	time, err := ParseTime(timeString)
	if err != nil {
		t.Errorf("Shouldn't error with given string: %v", timeString)
	}

	if time.Hour != 14 {
		t.Errorf("Not the hour, which is expected, got: %v, wanted: %v", time.Hour, 14)
	}
	if time.Minute != 7 {
		t.Errorf("Not the minute, which is expected, got: %v, wanted: %v", time.Minute, 7)
	}
	if time.Second != 33 {
		t.Errorf("Not the second, which is expected, got: %v, wanted: %v", time.Second, 7)
	}

	fmt.Println(time)
}