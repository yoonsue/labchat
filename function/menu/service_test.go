package menu

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetSchool(t *testing.T) {
	// var s *service
	testCases := []struct {
		url          string
		expectedRest string
	}{
		{
			"http://www.hanyang.ac.kr/web/www/-254",
			"교직원식당",
		},
		{
			"http://www.hanyang.ac.kr/web/www/-255",
			"학생식당",
		},
		{
			"http://www.hanyang.ac.kr/web/www/-258",
			"창업보육센터",
		},
	}
	for _, c := range testCases {
		///// TO BE CHANGED : scrapMenu --> GetSchool
		gotMenuRest, _ := scrapMenu(c.url)

		if c.expectedRest != gotMenuRest {
			t.Errorf("expected %s, got %s", c.expectedRest, gotMenuRest)
		}
	}
}
func TestScrapMenu(t *testing.T) {
	testCases := []struct {
		url string
	}{
		{
			"http://www.hanyang.ac.kr/web/www/-254",
		},
		{
			"http://www.hanyang.ac.kr/web/www/-255",
		},
		{
			"http://www.hanyang.ac.kr/web/www/-258",
		},
	}

	// TO BE IMPLEMENTED: check it is OK to fuction seperate to scrapMenu
	for _, c := range testCases {
		_, err := goquery.NewDocument(c.url)
		if err != nil {
			t.Errorf("failed to get URL")
		}
	}
}
