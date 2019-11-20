package tidy

import "testing"

import "time"

// JavaScript Date object `Date.parse` method return string format is `RFC2822` or `ISO8601`
func Test_TimeParse(t *testing.T) {
	s := "2019-11-07T07:32:37.375Z"
	origin, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Errorf("parse time error %v", err)
		return
	}

	if ss := origin.Local().Format("2006-01-02 15:04:05"); ss != "2019-11-07 15:32:37" {
		t.Errorf("location time format error %v", err)
		return
	}
}
