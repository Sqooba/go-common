package time

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CurrentTimestampInMS returns the current epoch time in milliseconds instead of nanos.
func CurrentTimestampInMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// FormatTimestampInMS formats a timestamp provided in milliseconds.
func FormatTimestampInMS(format string, ms int64) string {
	return time.Unix(0, ms*int64(time.Millisecond)).Format(format)
}

var durationPattern = regexp.MustCompile(`(-?\d+)\s*([a-zA-Z]*)`)

// ParseDuration extracts a duration in days from the passed string.
// supports expressions like 1s (or sec, second or seconds),
// 2m (or mi, min or minutes), 3h (hours), 4d (or days), 5mo (or months),
// 6y (or years), with default (i.e. no unit) to seconds.
// Returns Duration of 0 in case of parsing error.
func ParseDuration(duration string) time.Duration {

	matches := durationPattern.FindStringSubmatch(duration)

	if len(matches) < 3 || matches[1] == "" {
		log.Printf("Faulty duration string: %s", duration)
		return time.Duration(0)
	}
	dStr := matches[1]
	dInt, err := strconv.ParseInt(dStr, 10, 64)
	if err != nil {
		log.Println(err)
		return time.Duration(0)
	}

	u := matches[2]

	switch strings.ToLower(u) {

	case "", "s", "sec", "second", "seconds":

	case "m", "mi", "min", "minute", "minutes":
		dInt *= 60
	case "h", "hour", "hours":
		dInt *= 60 * 60
	case "d", "day", "days":
		dInt *= 60 * 60 * 24
	case "w", "week", "weeks":
		dInt *= 60 * 60 * 24 * 7
	case "mo", "mon", "month", "months":
		dInt *= 60 * 60 * 24 * 30
	case "y", "year", "years":
		dInt *= 60 * 60 * 24 * 30 * 365
	default:
		log.Fatal("unrecognized duration unit: " + u)
	}
	return time.Duration(dInt) * time.Second
}

// Unix epoch (or Unix time or POSIX time or Unix timestamp)  1 year (365.24 days)
const infinity float64 = 31556926 * 1000

// timeLapse condition struct
type timeLapse struct {
	// Time stamp threshold to handle the time lap condition
	Threshold float64
	Divider   float64
	// Handler function which determines the time lapse based on the condition
	TimePeriod string
}

var timeLapses = []timeLapse{
	{Threshold: 60, Divider: 1, TimePeriod: "second"},
	{Threshold: 3600, Divider: 60, TimePeriod: "minute"},
	{Threshold: 86400, Divider: 3600, TimePeriod: "hour"},
	{Threshold: 604800, Divider: 86400, TimePeriod: "day"},
	{Threshold: 2592000, Divider: 604800, TimePeriod: "week"},
	{Threshold: 31536000, Divider: 2592000, TimePeriod: "month"},
	{Threshold: infinity, Divider: 31536000, TimePeriod: "year"},
}

// FormatDuration returns a string describing how long it has been since
// the time argument passed int
func FormatDuration(d time.Duration) (timeSince string) {

	for _, formatter := range timeLapses {
		if d.Seconds() < formatter.Threshold {
			n := d.Seconds() / formatter.Divider
			nStr := strconv.FormatFloat(n, 'f', 0, 64)
			if n >= 1.5 {
				return "" + nStr + " " + formatter.TimePeriod + "s"
			}
			return "" + nStr + " " + formatter.TimePeriod
		}
	}
	return ""
}