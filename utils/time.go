package utils

import "time"

// TimeNowString returns current time in string
func TimeNowString() string {
	return time.Now().Format(time.RFC3339)
}
