package utils

import "time"

func AddDays(t time.Time, days int) time.Time {
	return t.Add(time.Hour * 24 * time.Duration(days))
}

