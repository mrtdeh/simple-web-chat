package models

import (
	"database/sql/driver"
)

type Gesture string

const (
	SMILE Gesture = "SMILE"
	HAPPY Gesture = "HAPPY"
)

func (ct *Gesture) Scan(value interface{}) error {
	*ct = Gesture(value.([]byte))
	return nil
}

func (ct Gesture) Value() (driver.Value, error) {
	return string(ct), nil
}
