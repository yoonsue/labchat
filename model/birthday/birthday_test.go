package birthday

import (
	"strconv"
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {

	currentTime := time.Now().Local()
	currentString := currentTime.Format("2006-01-02")
	currentYear, _ := strconv.Atoi(currentString[1:4])

	testCases := []struct {
		birth    Birthday
		expected int
	}{
		{
			Birthday{Name: "name1", DateOfBirth: 000101},
			(currentYear + 1),
		},
		{
			Birthday{Name: "name1", DateOfBirth: 990101},
			(100 + currentYear - 99 + 1),
		},
	}
	for _, c := range testCases {
		gotAge := c.birth.GetAge()
		if c.expected != gotAge {
			t.Errorf("expected %d, got %d", c.expected, gotAge)
		}
	}
}
func TestGetBirth(t *testing.T) {
	testCases := []struct {
		birth    Birthday
		expected string
	}{
		{
			Birthday{Name: "name01", DateOfBirth: 991201},
			"1999년 12월 01일",
		},
	}
	for _, c := range testCases {
		gotBirth := c.birth.GetBirth()
		if c.expected != gotBirth {
			t.Errorf("expected %s, got %s", c.expected, gotBirth)
		}
	}
}
