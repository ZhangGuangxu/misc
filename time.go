package misc

import "time"

// SleepMSec sleeps ms milliseconds
func SleepMSec(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

// GetUnixMSec returns timestamp accurate to millisecond
func GetUnixMSec() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
