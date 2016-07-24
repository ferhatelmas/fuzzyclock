package fuzzyclock

import (
	"fmt"
	"time"
)

var hours = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

var formats = []string{
	"%s o'clock",
	"five past %s",
	"ten past %s",
	"quarter past %s",
	"twenty past %s",
	"twenty five past %s",
	"half past %s",
	"twenty five to %s",
	"twenty to %s",
	"quarter to %s",
	"ten to %s",
	"five to %s",
}

func hourName(h int) string {
	if h < 1 {
		h = 12
	}
	return hours[(h-1)%12]
}

// Time returns given time(hour and minute) in human readable format (fuzzy)
func Time(h, m int) string {
	var i int
	if m > 2 {
		i = ((m - 3) / 5) + 1
	}
	if i > 6 {
		h++
	}
	return fmt.Sprintf(formats[i%12], hourName(h))
}

// Now returns currenct local time in fuzzy format
func Now() string {
	t := time.Now()
	return Time(t.Hour(), t.Minute())
}
