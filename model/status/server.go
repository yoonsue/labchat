package status

import (
	"strconv"
	"time"
)

// Server = VO
type Server struct {
	Temperature Temperature
	TimeStamp   time.Time
}

const timeFormat string = "2006-01-02 15:04:05 "

// String returns the formatted string of server status information.
func (s Server) String() string {
	str := "TIME : " + s.TimeStamp.Format(timeFormat)
	str = str + "\n"
	str = "TEMP : " + s.Temperature.String()
	return str
}

// Temperature = VO
type Temperature float32

// String returns the temperature with human readable format.
func (t Temperature) String() string {
	return strconv.FormatFloat(float64(t), 'g', -1, 64) + " C"
}
