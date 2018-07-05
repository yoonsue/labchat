package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

// URL1: http://www.hanyang.ac.kr/web/www/-255	학생식당
// //*[@id="messhall1"]/div/div/div/div/ul/li/a/h3
// //*[@id="messhall1"]/div/div/div/div/ul/li/a/img
// URL2: http://www.hanyang.ac.kr/web/www/-256	창의인재원식당
// URL3: http://www.hanyang.ac.kr/web/www/-258	창업보육센터

// GetMenu gets specific HTML body
func GetMenu(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get response from HTML URL"))
	}

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read HTML body"))
	}
	// root, err := html.Parse(resp.Body)
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "failed to parse HTML body"))
	// }
	// fmt.Printf("\n\nPARSED:%s\n", root)
	//element, ok := getElementById("login_field", root)
	return html
}

// GetMenuColly gets menu from HTML body
func GetMenuColly(url string) string {
	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		fmt.Println("Found: ", title)
	})
	return "hello"
}

func postScrap(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get URL"))
	}

	doc.Find("#messhall1 div div div div ul li").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		menu := item.Find("h3")
		fmt.Printf("Post #%d: %s - %s\n", index, title, menu)

	})
}
