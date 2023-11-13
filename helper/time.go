package helper

import (
	"github.com/golang-module/carbon/v2"
)

func GetUnixTimeInMs() int64 {
	return carbon.Now("Asia/Kolkata").TimestampMilli()
}

func GetUnixTimeInNanoSecond() int64 {
	return carbon.Now("Asia/Kolkata").TimestampNano()
}
