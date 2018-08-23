package status

import (
	"testing"

	model "github.com/yoonsue/labchat/model/status"
)

func TestServerCheck(t *testing.T) {
	s := &model.Server{
		Temperature: 100,
	}

	gotS := model.NewServer()
	gotS.Temperature = 100

	if s.Temperature != gotS.Temperature {
		t.Errorf("expected %s, got %s", s.Temperature, gotS.Temperature)
	}
}

// Cant get access in code coverage /sys.
func TestGetTemp(t *testing.T) {
	t.Skip("skipping test there is no temp file in code coverage")
}
