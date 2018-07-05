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

// Scrap gets element that you want
func Scrap(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get URL"))
	}
	// //*[@id="content"]/div[2]/div[1]/div/h3
	// //*[@id="yui_patched_v3_11_0_1_1530810633127_193"]/div[1]/div/h3
	// //*[@id="p_p_id_56_INSTANCE_3N8IwmXJbuCj_"]/div/div/div[1]/div/div[2]/h4/strong
	// title := doc.Find("#content div[2] div div h3").Text()
	// // titleText := title.Text()
	// fmt.Printf("title: %s\n", title)

	doc.Find("#messhall1 div div div div ul li").Each(func(index int, item *goquery.Selection) {
		doc.Find("h3")
		menu := item.Find("h3")
		menuText := menu.Text()
		fmt.Printf("Post #%d: %s \n", index, menuText)
		// return ("Post #" + string(index) + ": " + title + " - " + menuText + "\n")
	})

	return "_"
}
