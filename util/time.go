package util

import "time"

type DateTime time.Time

func NewDateTime(convert *time.Time) *DateTime {
	if convert != nil {
		x := DateTime(*convert)
		return &x
	}
	return nil
}
