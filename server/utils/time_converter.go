package utils

import "time"

func StringToTime(date string) (time.Time, error) {
	converted, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	} else {
		return converted, nil
	}
}
