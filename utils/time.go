package utils

import "time"

// TimeNowString returns current time in string
func TimeNowString() string {
	return time.Now().Format(time.RFC3339)
}

// Time2Timestamp convert Time to (10, 13) bit unix timestamp
func Time2Timestamp(t time.Time, bit int) int64 {
	switch bit {
	case 10:
		return t.UnixNano() / int64(time.Second)
	case 13:
		return t.UnixNano() / int64(time.Millisecond)
	default:
		return t.UnixNano() / int64(time.Millisecond)
	}
}

// NowTimestamp ..
func NowTimestamp(bit int) int64 {
	return Time2Timestamp(time.Now(), bit)
}
