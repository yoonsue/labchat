package server

import (
	"strings"
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

func TestMsgFor(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{
			strings.Fields("lab status"),
			"TIME : ", //TO BE IMPLEMENTED
		},
		{
			strings.Fields("lab menu"),
			"\n==교직원식당==\n", //TO BE IMPLEMENTED
		},
		{
			strings.Fields("shit"),
			"shit....????",
		},
	}

	for _, c := range testCases {
		result := msgFor(c.input)
		if !strings.HasPrefix(result, c.expected) {
			t.Errorf("start with '%s' on input %s, got %s", c.expected, c.input, result)
		}
	}
}
