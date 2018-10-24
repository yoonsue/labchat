package phone

import (
	"testing"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		depart   Department
		expected string
	}{
		{
			"department1",
			"department1",
		},
	}
	for _, c := range testCases {
		gotDepart := c.depart.ToString()
		if c.expected != gotDepart {
			t.Errorf("expected %s, got %s", c.expected, gotDepart)
		}
	}
}
