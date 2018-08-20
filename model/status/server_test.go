package status

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s := &Server{
		timeStamp: time.Now().Format(timeFormat),
	}
	gotServer := NewServer()
	if s.timeStamp != gotServer.timeStamp {
		t.Errorf("expected %s, got %s", s.timeStamp, gotServer.timeStamp)
	}
}
