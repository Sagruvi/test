package entity

import (
	"time"
)

type Event struct {
	EventType int       `json:"eventID" db:"eventID"`
	UserID    int       `json:"userID" db:"userID"`
	EventTime time.Time `json:"eventTime"`
	Payload   int       `json:"payload"`
}
type Request struct {
	EventType int       `json:"eventID"`
	After     time.Time `json:"upper limit"`
	Before    time.Time `json:"lower limit"`
}
