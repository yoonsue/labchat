package function

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model"
)

// URL1: http://www.hanyang.ac.kr/web/www/-255	학생식당
// //*[@id="messhall1"]/div/div/div/div/ul/li/a/h3
// //*[@id="messhall1"]/div/div/div/div/ul/li/a/img
// URL2: http://www.hanyang.ac.kr/web/www/-256	창의인재원식당
// URL3: http://www.hanyang.ac.kr/web/www/-258	창업보육센터

// MenuGet assigns menu to Menu model
func MenuGet(url string) model.Menu {
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
		// doc.Find("a img")
		menu := item.Find("a img")
		menuTitle, _ := menu.Attr("alt")
		fmt.Printf("Post #%d: %s \n", index, menuTitle)
		menuText += menuTitle + "\n"
	})
	return menuText
}
