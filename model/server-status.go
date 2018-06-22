package model

import (
	"time"
)

// ServerStatus = VO
type ServerStatus struct {
	Temperature Temperature
	TimeStamp   time.Time
}

// Temperature = VO
type Temperature float32
