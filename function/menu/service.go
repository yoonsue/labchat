package menu

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	model "github.com/yoonsue/labchat/model/menu"
)

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
	// menu.Title := scrapMenu(url)
	menu.Menu = scrapMenu(url)
	return menu
}

// scrapMenu gets menu from the URL
func scrapMenu(url string) string {
	var menuText = ""
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get URL"))
	}
	doc.Find("#messhall1 div div div div ul li").Each(func(index int, item *goquery.Selection) {
		menu := item.Find("a img")
		menuTitle, _ := menu.Attr("alt")
		menuText += menuTitle + "\n"
	})
	return menuText
}

func NewService(r model.Repository) Service {
	return &service{
		menus: r,
	}
}
