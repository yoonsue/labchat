package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	fmt.Printf("\n\nORIGINAL:\n%s\n", html)
	// root, err := html.Parse(resp.Body)
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "failed to parse HTML body"))
	// }
	// fmt.Printf("\n\nPARSED:%s\n", root)
	//element, ok := getElementById("login_field", root)
	return html
}
