package time

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
)

const (
	oneNanoSecondInSecond = 1000000000
)

func TestParseDuration(t *testing.T) {
	assert.Equal(t, time.Duration(2*oneNanoSecondInSecond), ParseDuration("2"))
	assert.Equal(t, time.Duration(2*oneNanoSecondInSecond), ParseDuration("2.9"))
	assert.Equal(t, time.Duration(-oneNanoSecondInSecond), ParseDuration("-1"))
	assert.Equal(t, time.Duration(-2*oneNanoSecondInSecond), ParseDuration("-2"))
	assert.Equal(t, time.Duration(1*60*oneNanoSecondInSecond), ParseDuration("1m"))
	assert.Equal(t, time.Duration(2*60*oneNanoSecondInSecond), ParseDuration("2min"))
	assert.Equal(t, time.Duration(4*60*oneNanoSecondInSecond), ParseDuration("4minutes"))
	assert.Equal(t, time.Duration(2*60*60*24*oneNanoSecondInSecond), ParseDuration("2d"))
	assert.Equal(t, time.Duration(3*60*60*24*oneNanoSecondInSecond), ParseDuration("3day"))
	assert.Equal(t, time.Duration(6*60*60*24*oneNanoSecondInSecond), ParseDuration("6days"))
	assert.Equal(t, time.Duration(1*7*60*60*24*oneNanoSecondInSecond), ParseDuration("1week"))
	assert.Equal(t, time.Duration(3*7*60*60*24*oneNanoSecondInSecond), ParseDuration("3weeks"))
	assert.Equal(t, time.Duration(1*30*60*60*24*oneNanoSecondInSecond), ParseDuration("1month"))
	assert.Equal(t, time.Duration(10*30*60*60*24*oneNanoSecondInSecond), ParseDuration("10mo"))
	assert.Equal(t, time.Duration(24*30*60*60*24*oneNanoSecondInSecond), ParseDuration("24mon"))
	assert.Equal(t, time.Duration(36*30*60*60*24*oneNanoSecondInSecond), ParseDuration("36months"))
	assert.Equal(t, time.Duration(0), ParseDuration("invalidString"))
}

func TestFormatDuration(t *testing.T) {
	assert.Equal(t, "0 second", FormatDuration(time.Duration(oneNanoSecondInSecond/10)))
	assert.Equal(t, "1 second", FormatDuration(time.Duration(oneNanoSecondInSecond)))
	assert.Equal(t, "2 seconds", FormatDuration(time.Duration(2*oneNanoSecondInSecond)))
	assert.Equal(t, "10 seconds", FormatDuration(time.Duration(10*oneNanoSecondInSecond)))
	assert.Equal(t, "59 seconds", FormatDuration(time.Duration(59*oneNanoSecondInSecond)))
	assert.Equal(t, "1 minute", FormatDuration(time.Duration(60*oneNanoSecondInSecond)))
	assert.Equal(t, "2 minutes", FormatDuration(time.Duration(2*60*oneNanoSecondInSecond)))
	assert.Equal(t, "59 minutes", FormatDuration(time.Duration(59*60*oneNanoSecondInSecond)))
	assert.Equal(t, "1 hour", FormatDuration(time.Duration(60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "2 hours", FormatDuration(time.Duration(2*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "23 hours", FormatDuration(time.Duration(23*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "1 day", FormatDuration(time.Duration(26*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "2 days", FormatDuration(time.Duration(45*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "2 days", FormatDuration(time.Duration(36*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "2 weeks", FormatDuration(time.Duration(15*24*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "1 month", FormatDuration(time.Duration(30*24*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "1 month", FormatDuration(time.Duration(35*24*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "2 months", FormatDuration(time.Duration(60*24*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "1 year", FormatDuration(time.Duration(365*24*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "10 years", FormatDuration(time.Duration(3650*24*60*60*oneNanoSecondInSecond)))
	assert.Equal(t, "100 years", FormatDuration(time.Duration(36500*24*60*60*oneNanoSecondInSecond)))
}
