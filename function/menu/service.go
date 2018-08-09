package menu

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	model "github.com/yoonsue/labchat/model/menu"
)

// Service declares the methods that menu service provides.
type Service interface {
	GetSchool(url string) model.Menu
}

type service struct {
	menus model.Repository
}

// URL1: http://www.hanyang.ac.kr/web/www/-255	학생식당
// //*[@id="messhall1"]/div/div/div/div/ul/li/a/img
// URL2: http://www.hanyang.ac.kr/web/www/-256	창의인재원식당
// URL3: http://www.hanyang.ac.kr/web/www/-258	창업보육센터

// GetSchool assigns menu to Menu model
func (s *service) GetSchool(url string) model.Menu {
	menu := model.Menu{}
	restText, menuText := scrapMenu(url)
	menu.Restaurant = model.Restaurant(restText)
	menu.TodayMenu = model.TodayMenu(menuText)
	s.menus.Store(menu)
	resMenu, err := s.menus.Find(menu.Restaurant)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to find menu in menuMap"))
	}
	return resMenu
}

// scrapMenu gets menu from the URL
func scrapMenu(url string) (string, string) {
	menuText := ""
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get URL"))
	}

	restTitle := doc.Find("#content div.contents-box div.title-top div h3").Text()
	doc.Find("#messhall1 div div div div ul li").Each(func(index int, item *goquery.Selection) {
		menu := item.Find("a img")
		menuTitle, _ := menu.Attr("alt")
		menuText += menuTitle + "\n"
	})
	return restTitle, menuText
}

// #content > div.contents-box.container > div.contents-title-arrow.title-top.sub > div > h3

// NewService return struct which provides Service interface
func NewService(r model.Repository) Service {
	return &service{
		menus: r,
	}
}
