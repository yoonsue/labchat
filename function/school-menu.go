package function

import "golang.org/x/net/html"

// URL1: http://www.hanyang.ac.kr/web/www/-255	학생식당
// //*[@id="messhall1"]/div/div/div/div/ul/li/a/h3
// //*[@id="messhall1"]/div/div/div/div/ul/li/a/img
// URL2: http://www.hanyang.ac.kr/web/www/-256	창의인재원식당
// URL3: http://www.hanyang.ac.kr/web/www/-258	창업보육센터

func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		href = a.Val
		ok = true
	}
	return
}

func crawl(url string) {

}
