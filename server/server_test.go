package server

import (
	"testing"
)

func TestMessageKey(t *testing.T) {
	testCases := []struct {
		key      string
		expected string
	}{
		{
			"hi",
			"hello", // TestCase 1
		},
		{
			"hello",
			"hi",
		},
		{
			"name",
			"LABchat",
		},
		{
			"no",
			"none",
		},
	}

	for _, c := range testCases {
		result := messageKey(c.key)
		if result != c.expected {
			t.Errorf("expected %s for key %s, got %s", c.expected, c.key, result)
		}
	}
}
