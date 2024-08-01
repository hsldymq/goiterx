package goiterx

import (
	"time"

	"github.com/hsldymq/goiter"
)

// RangeTime is similar to RangeStep, but it is specifically used for iterating over time, and it can iterate time forward or backward.
// The interval parameter is its step size, which can be any positive duration.
// Unlike the half-open interval represented by the start and end parameters of RangeStep, the from and to parameters of RangeTime represent a closed interval.
// For example:
//
//	from := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
//	to := time.Date(2021, 1, 1, 5, 0, 0, 0, time.UTC)
//	for t := range RangeTime(from, to, time.Hour) {
//	    fmt.Println(t.Format(time.TimeOnly))
//	}
//	the above code will print:
//		00:00:00
//		01:00:00
//		02:00:00
//		03:00:00
//		04:00:00
//		05:00:00
func RangeTime(from time.Time, to time.Time, interval time.Duration) goiter.Iterator[time.Time] {
	if interval <= 0 {
		return goiter.Empty[time.Time]()
	}

	return func(yield func(time.Time) bool) {
		if from.Before(to) || from.Equal(to) {
			t := from
			for t.Before(to) || t.Equal(to) {
				if !yield(t) {
					return
				}
				t = t.Add(interval)
			}
		} else {
			t := from
			for t.After(to) || t.Equal(to) {
				if !yield(t) {
					return
				}
				t = t.Add(-interval)
			}
		}
	}
}
