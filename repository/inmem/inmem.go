package inmem

import (
	"sync"

	"github.com/yoonsue/labchat/model/menu"
	"github.com/yoonsue/labchat/model/phone"
)

// MenuRepository struct definition.
type MenuRepository struct {
	mtx     sync.RWMutex
	menuMap map[menu.Restaurant]*menu.Menu
}

// NewMenuRepository return a new instance of in-memory menu repository.
func NewMenuRepository() menu.Repository {
	return &MenuRepository{
		menuMap: make(map[menu.Restaurant]*menu.Menu),
	}
}

// Store saves menu model in memory.
func (r *MenuRepository) Store(target *menu.Menu) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.menuMap[target.Restaurant] = target
	return nil
}

// Find returns today's menus that match with the given restaurant.
func (r *MenuRepository) Find(rest menu.Restaurant) (*menu.Menu, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if menu, exists := r.menuMap[rest]; exists {
		return menu, nil
	}
	return nil, nil
}

// FindAll returns all menus that were stored in memory.
func (r *MenuRepository) FindAll() []*menu.Menu {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	c := make([]*menu.Menu, 0, len(r.menuMap))
	for _, val := range r.menuMap {
		c = append(c, val)
	}
	return c
}

// // menuURLMap
// var menuURLMap = map[string]string{
// 	"교직원식당":   "http://www.hanyang.ac.kr/web/www/-254",
// 	"학생식당":    "http://www.hanyang.ac.kr/web/www/-255",
// 	"창업보육센터":  "http://www.hanyang.ac.kr/web/www/-258",
// 	"창의인재원식당": "http://www.hanyang.ac.kr/web/www/-256",
// }

// var menuMap = map[string]string{
// 	"창의인재원":  "",
// 	"학생식당":   "",
// 	"교직원식당":  "",
// 	"창업보육센터": "",
// }

// func allMenu() string {
// 	str := ""
// 	for key := range menuMap {
// 		str += key
// 		str += "\t"
// 		str += menuMap[key]
// 		str += "\n"
// 		fmt.Printf("key: %s\tvalue: %s", key, menuMap[key])
// 	}
// 	return str
// }

// PhoneRepository struct definition.
type PhoneRepository struct {
	mtx      sync.RWMutex
	phoneMap map[phone.Department]*phone.Phone
}

// NewPhoneRepository return a new instance of in-memory phone repository.
func NewPhoneRepository() phone.Repository {
	return &PhoneRepository{
		phoneMap: make(map[phone.Department]*phone.Phone),
	}
}

// Store saves menu model in memory.
func (r *PhoneRepository) Store(target *phone.Phone) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.phoneMap[target.Department] = target
	return nil
}

// Find returns today's menus that match with the given restaurant.
func (r *PhoneRepository) Find(d phone.Department) (*phone.Phone, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if phone, exists := r.phoneMap[d]; exists {
		return phone, nil
	}
	return nil, nil
}
