package helper

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
)

func GetUnixTimeInMs() int64 {

	fmt.Println(carbon.Now().ToTimestampMilliStruct())
	loc, _ := time.LoadLocation("Asia/Kolkata")

	return time.Now().In(loc).UnixMilli()
}

func GetUnixTimeInNanoSecond() int64 {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc).UnixNano()
}
