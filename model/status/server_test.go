package status

import (
	"testing"
)

func TestTime(t *testing.T) {
	t.Skip("No need to test...")
}

// func TestNewServer(t *testing.T) {
// 	s := &Server{
// 		timeStamp: time.Now().Format(timeFormat),
// 	}
// 	gotServer := NewServer()
// 	if s.timeStamp != gotServer.timeStamp {
// 		t.Errorf("expected %s, got %s", s.timeStamp, gotServer.timeStamp)
// 	}
// }
func TestString(t *testing.T) {
	testCases := []struct {
		temp     Temperature
		expected string
	}{
		{
			3.0,
			"3 C",
		},
		{
			3.9,
			"3.9000000953674316 C",
		},
		{
			-1,
			"-1 C",
		},
	}
	for _, c := range testCases {
		gotTemp := c.temp.String()
		if c.expected != gotTemp {
			t.Errorf("expected %s, got %s", c.expected, gotTemp)
		}
	}
}
