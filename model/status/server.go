package status

import (
	"strconv"
	"time"
)

const timeFormat string = "2006-01-02 15:04:05 "

// Server = VO
type Server struct {
	Temperature Temperature
	timeStamp   string
}

// Time returns the time of checking server request.
func (s Server) Time() string {
	return s.timeStamp
}

// NewServer fixes the timestamp when the request is arrived.
func NewServer() *Server {
	return &Server{
		timeStamp: time.Now().Format(timeFormat),
	}
}

// Temperature = VO
type Temperature float32

// String returns the temperature with human readable format.
func (t Temperature) String() string {
	return strconv.FormatFloat(float64(t), 'g', -1, 64) + " C"
}
