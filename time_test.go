package goiterx

import (
	"slices"
	"testing"
	"time"
)

func TestRangeTime(t *testing.T) {
	from, _ := time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ := time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual := make([]string, 0)
	for t := range RangeTime(from, to, 24*time.Hour) {
		actual = append(actual, t.Format(time.DateOnly))
	}
	expect := []string{"2024-01-01", "2024-01-02", "2024-01-03", "2024-01-04", "2024-01-05"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:10:00")
	to, _ = time.Parse(time.DateTime, "2024-01-01 03:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{"2024-01-01 00:10:00", "2024-01-01 01:10:00", "2024-01-01 02:10:00"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ = time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual = make([]string, 0)
	i := 0
	for t := range RangeTime(from, to, 24*time.Hour) {
		actual = append(actual, t.Format(time.DateOnly))
		i++
		if i >= 3 {
			break
		}
	}
	expect = []string{"2024-01-01", "2024-01-02", "2024-01-03"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 03:10:00")
	to, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{"2024-01-01 03:10:00", "2024-01-01 02:10:00", "2024-01-01 01:10:00", "2024-01-01 00:10:00"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 03:10:00")
	to, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	actual = make([]string, 0)
	i = 0
	for t := range RangeTime(from, to, time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
		i++
		if i >= 3 {
			break
		}
	}
	expect = []string{"2024-01-01 03:10:00", "2024-01-01 02:10:00", "2024-01-01 01:10:00"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ = time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, time.Duration(0)) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ = time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, -time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}
}
