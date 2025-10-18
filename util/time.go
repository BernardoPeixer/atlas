package util

import (
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

func NewDateTime(convert *time.Time) *DateTime {
	if convert != nil {
		x := DateTime(*convert)
		return &x
	}
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	date := time.Time(d).Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf("%q", date)), nil
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	date := strings.Trim(string(b), `"`)
	if date == "null" {
		*d = DateTime(time.Time{})
		return nil
	}
	tmp, err := time.Parse("2006-01-02 15:04:05", date)
	if err == nil {
		*d = DateTime(tmp)
	}
	return err
}

func (d *DateTime) Time() *time.Time {
	if d == nil {
		return nil
	}
	t := time.Time(*d)
	return &t
}
