package helper

import "time"

func GetUnixTimeInMs() int64 {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc).UnixMilli()
}

func GetUnixTimeInNanoSecond() int64 {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc).UnixNano()
}
