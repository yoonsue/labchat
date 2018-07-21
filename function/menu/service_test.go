package menu

import (
	"strings"
	"testing"
)

func TestScrapMenu(t *testing.T) {
	testCases := []struct {
		url      string
		expected string
	}{
		{
			"http://www.hanyang.ac.kr/web/www/-254",
			"[중식A]",
		},
		{
			"http://www.hanyang.ac.kr/web/www/-255",
			"[특식]",
		},
		{
			"http://www.hanyang.ac.kr/web/www/-258",
			"[한식]",
		},
	}

	for _, c := range testCases {
		result := scrapMenu(c.url)
		if !strings.HasPrefix(result, c.expected) {
			t.Errorf("start with '%s' on url %s, got '%s'", c.expected, c.url, result)
		}
	}
}
