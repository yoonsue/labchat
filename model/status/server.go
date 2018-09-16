package status

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const timeFormat string = "2006-01-02 15:04:05 "

// Server = VO
type Server struct {
	Temperature Temperature
	BootTime    time.Time
	Uptime      time.Duration
}

// Time returns the time of checking server request.
func (s Server) Time() string {
	return time.Now().Format(timeFormat)
}

// FmtDuration returns time string.
func FmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	day := d / (24 * time.Hour)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("day%03d, %02d:%02d", day, h, m)
}

// NewServer fixes the timestamp when the request is arrived.
func NewServer() Server {
	log.Println("inside newserver")
	return Server{
		BootTime: time.Now().Round(0).Add(time.Second),
	}
}

// Temperature = VO
type Temperature float32

// String returns the temperature with human readable format.
func (t Temperature) String() string {
	return strconv.FormatFloat(float64(t), 'g', -1, 64) + " C"
}
