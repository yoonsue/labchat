package inmem

import (
	"strings"
	"sync"

	"github.com/yoonsue/labchat/model/birthday"

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

// Store saves phone model in memory.
func (r *PhoneRepository) Store(target *phone.Phone) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.phoneMap[target.Department] = target
	return nil
}

// Find returns phone that match with the given restaurant.
func (r *PhoneRepository) Find(d phone.Department) ([]*phone.Phone, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	var phoneList []*phone.Phone
	for key, phone := range r.phoneMap {
		if strings.Contains(key.ToString(), d.ToString()) {
			phoneList = append(phoneList, phone)
		}
	}
	return phoneList, nil

}

// BirthdayRepository struct definition.
type BirthdayRepository struct {
	mtx         sync.RWMutex
	birthdayMap map[string]*birthday.Birthday
}

// NewBirthdayRepository return a new instance of in-memory phone repository.
func NewBirthdayRepository() birthday.Repository {
	return &BirthdayRepository{
		birthdayMap: make(map[string]*birthday.Birthday),
	}
}

// Store saves phone model in memory.
func (r *BirthdayRepository) Store(target *birthday.Birthday) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.birthdayMap[target.Name] = target
	return nil
}

// Find returns phone that match with the given restaurant.
func (r *BirthdayRepository) Find(name string) (*birthday.Birthday, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if birthday, exists := r.birthdayMap[name]; exists {
		return birthday, nil
	}
	return nil, nil
}
