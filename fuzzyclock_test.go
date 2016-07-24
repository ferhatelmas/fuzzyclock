package fuzzyclock

import (
	"fmt"
	"testing"
)

func TestHourName(t *testing.T) {
	for i := 1; i < 13; i++ {
		if hourName(i) != hours[i-1] {
			t.Errorf("%d: expected %s but got %s", i, hours[i-1], hourName(i))
		}
	}
}

func TestHourNameZero(t *testing.T) {
	if hourName(0) != "twelve" {
		t.Errorf("hour name for 0 must be twelve, got %s", hourName(0))
	}
}

func gen(h, m, s, e int, format string, onlyMin bool, t *testing.T) {
	ch, cm := h, m
	if onlyMin && s >= 33 {
		h++
	}
	for ; s < e; s++ {
		var expected string
		if onlyMin {
			cm = s
			expected = fmt.Sprintf(format, hourName(h))
		} else {
			ch = s
			expected = fmt.Sprintf(format, hourName(s))
		}
		if Time(ch, cm) != expected {
			t.Errorf("%d: expected %s, got %s", s, expected, Time(ch, cm))
		}
	}
}

func TestOClock(t *testing.T) {
	for i := 0; i < 3; i++ {
		gen(1, i, 0, 13, "%s o'clock", false, t)
	}
}
func TestFivePast(t *testing.T) {
	gen(1, 0, 3, 8, "five past %s", true, t)
}
func TestTenPast(t *testing.T) {
	gen(1, 0, 8, 13, "ten past %s", true, t)
}
func TestQuarterPast(t *testing.T) {
	gen(1, 0, 13, 18, "quarter past %s", true, t)
}
func TestTwentyPast(t *testing.T) {
	gen(1, 0, 18, 23, "twenty past %s", true, t)
}
func TestTwentyFivePast(t *testing.T) {
	gen(1, 0, 23, 28, "twenty five past %s", true, t)
}
func TestHalfPast(t *testing.T) {
	gen(1, 0, 28, 33, "half past %s", true, t)
}
func TestTwentyFiveTo(t *testing.T) {
	gen(1, 0, 33, 38, "twenty five to %s", true, t)
}
func TestTwentyTo(t *testing.T) {
	gen(1, 0, 38, 43, "twenty to %s", true, t)
}
func TestQuarterTo(t *testing.T) {
	gen(1, 0, 43, 48, "quarter to %s", true, t)
}
func TestTenTo(t *testing.T) {
	gen(1, 0, 48, 53, "ten to %s", true, t)
}
func TestFiveTo(t *testing.T) {
	gen(1, 0, 53, 58, "five to %s", true, t)
}
func Test5859OclockFiveTo(t *testing.T) {
	gen(1, 58, 58, 60, "%s o'clock", true, t)
}
