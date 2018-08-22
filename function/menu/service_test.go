package menu

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	model "github.com/yoonsue/labchat/model/menu"
	"github.com/yoonsue/labchat/repository/inmem"
)

func TestGetSchool(t *testing.T) {
	r := inmem.NewMenuRepository()
	s := NewService(r)
	// var s *service
	testCases := []struct {
		url          string
		expectedRest model.Restaurant
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
		gotMenu := s.GetSchool(c.url)

		if c.expectedRest != gotMenu.Restaurant {
			t.Errorf("expected %s, got %s", c.expectedRest, gotMenu.Restaurant)
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

func TestNewService(t *testing.T) {
	r := inmem.NewMenuRepository()
	s := &service{
		menus: r,
	}

	gotService := NewService(r)

	tmpURL := "http://www.hanyang.ac.kr/web/www/-254"
	if s.GetSchool(tmpURL).Restaurant != gotService.GetSchool(tmpURL).Restaurant {
		t.Errorf("expected %s, got %s", s.GetSchool(tmpURL).Restaurant, gotService.GetSchool(tmpURL).Restaurant)
	}
}
