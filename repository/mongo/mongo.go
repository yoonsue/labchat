package mongo

import (
	"github.com/yoonsue/labchat/model/menu"
)

// MenuRepository struct definition
type MenuRepository struct{}

// NewMenuRepository does several services according to MongoDB
func NewMenuRepository() menu.Repository {
	return &MenuRepository{}
}

// Store saves menu model in memory.
func (r *MenuRepository) Store(target menu.Menu) error {
	return nil
}

// Find returns today's menus that match with the given restaurant.
func (r *MenuRepository) Find(rest menu.Restaurant) (menu.Menu, error) {
	menu := menu.Menu{}
	return menu, nil
}
