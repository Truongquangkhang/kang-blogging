package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var locale = time.UTC
var utils_time_logger = *logrus.StandardLogger().WithField("logger", "utils-time")

func init() {
	// if `TIMEZONE` is not set, timezoneString will be "", and time.LoadLocation defaults "" to "UTC"
	envTimezone := os.Getenv("TIMEZONE")
	loc, err := time.LoadLocation(envTimezone)
	if err != nil { // TIMEZONE not recognized by time package
		utils_time_logger.Fatal(fmt.Sprintf("Unrecognized timezone value='%s'", envTimezone))
	}
	locale = loc
}

func GetServerNow() time.Time {
	return time.Now().In(locale)
}

func ConvertTimeToTimestamp(date string) int {
	format := "2006-01-02 15:04:05"

	t, err := time.Parse(format, date)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t.Unix())
		return int(t.Unix())
	}
	return 0
}
