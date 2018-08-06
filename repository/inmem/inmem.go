package inmem

import (
	"fmt"
	"sync"

	"github.com/yoonsue/labchat/model/menu"
)

// MenuRepository struct definition
type MenuRepository struct {
	mtx     sync.RWMutex
	menuMap map[menu.Restaurant]menu.TodayMenu
}

// NewMenuRepository does several services according to Go Map
func NewMenuRepository() menu.Repository {
	return &MenuRepository{
		menuMap: make(map[menu.Restaurant]menu.TodayMenu),
	}
}

// Store ..
func (r *MenuRepository) Store(target menu.Menu) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.menuMap[target.Restaurant] = target.TodayMenu
	return nil
}

// Find ..
func (r *MenuRepository) Find(rest menu.Restaurant) (menu.TodayMenu, error) {
	if val, exists := r.menuMap[rest]; exists {
		return val, nil
	}
	return "none", nil
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

func allMenu() string {
	str := ""
	for key := range menuMap {
		str += key
		str += "\t"
		str += menuMap[key]
		str += "\n"
		fmt.Printf("key: %s\tvalue: %s", key, menuMap[key])
	}
	return str
}
