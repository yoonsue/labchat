package inmem

import (
	"strconv"
	"strings"
	"sync"

	"github.com/yoonsue/labchat/model/birthday"
	"github.com/yoonsue/labchat/model/location"
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

// Clean the menu repository.
func (r *MenuRepository) Clean() error {
	m := r.menuMap
	for k := range m {
		delete(m, k)
	}
	return nil
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

// FindByNum returns phone that match with the given phonbe number.
func (r *PhoneRepository) FindByNum(phoneNum int) ([]*phone.Phone, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	var phoneList []*phone.Phone
	for _, phone := range r.phoneMap {
		if phone.Extension == strconv.Itoa(phoneNum) {
			phoneList = append(phoneList, phone)
		}
	}
	return phoneList, nil
}

// Clean the phone repository.
func (r *PhoneRepository) Clean() error {
	m := r.phoneMap
	for k := range m {
		delete(m, k)
	}
	return nil
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

// FindAll returns phone that match with the given restaurant.
func (r *BirthdayRepository) FindAll() ([]*birthday.Birthday, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	b := make([]*birthday.Birthday, 0, len(r.birthdayMap))
	for _, val := range r.birthdayMap {
		b = append(b, val)
	}
	return b, nil
}

// Clean the birthday repository.
func (r *BirthdayRepository) Clean() error {
	m := r.birthdayMap
	for k := range m {
		delete(m, k)
	}
	return nil
}

// LocationRepository struct definition.
type LocationRepository struct {
	mtx          sync.RWMutex
	locationList map[string]*location.Location
}

// NewLocationRepository return a new instance of in-memory location repository.
func NewLocationRepository() location.Repository {
	return &LocationRepository{
		locationList: make(map[string]*location.Location),
	}
}

// Store saves location model in memory.
func (r *LocationRepository) Store(target *location.Location) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.locationList[target.Name] = target
	return nil
}

// Find returns location that match with the given name.
func (r *LocationRepository) Find(name string) ([]*location.Location, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	var locationList []*location.Location
	for key, loc := range r.locationList {
		if strings.Contains(key, name) {
			locationList = append(locationList, loc)
		}
	}
	return locationList, nil
}

// Clean the location repository.
func (r *LocationRepository) Clean() error {
	m := r.locationList
	for k := range m {
		delete(m, k)
	}
	return nil
}
