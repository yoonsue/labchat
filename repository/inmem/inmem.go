package inmem

import "github.com/yoonsue/labchat/model/menu"

// MenuRepository struct definition
type MenuRepository struct {
	key   string
	value []string
}

// NewMenuRepository does several services according to Go Map
func NewMenuRepository() menu.Repository {
	return &MenuRepository{}
}

// menuURLMap
var menuURLMap = map[string]string{
	"교직원식당":   "http://www.hanyang.ac.kr/web/www/-254",
	"학생식당":    "http://www.hanyang.ac.kr/web/www/-255",
	"창업보육센터":  "http://www.hanyang.ac.kr/web/www/-258",
	"창의인재원식당": "http://www.hanyang.ac.kr/web/www/-256",
}

var menuMap = map[string]string{
	"창의인재원":  "",
	"학생식당":   "",
	"교직원식당":  "",
	"창업보육센터": "",
}

func menuSave(title string, menu string) {
	menuMap[title] = menu
}

func menuRead(title string) string {
	val, exists := menuMap[title]
	if !exists {
		println("No '", title, "' exists")
	}
	return val
}
