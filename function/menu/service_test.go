package menu

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
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

	// TO BE IMPLEMENTED: check it is OK to fuction seperate to scrapMenu
	for _, c := range testCases {
		_, err := goquery.NewDocument(c.url)
		if err != nil {
			t.Errorf("failed to get URL")
		}
		// if !strings.HasPrefix(result, c.expected) {
		// 	t.Errorf("start with '%s' on url %s, got '%s'", c.expected, c.url, result)
		// }
	}
}
