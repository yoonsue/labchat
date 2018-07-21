package status

import (
	"time"
)

// Server = VO
type Server struct {
	Temperature Temperature
	TimeStamp   time.Time
}

// Temperature = VO
type Temperature float32
