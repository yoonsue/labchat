package inmem

import (
	"github.com/yoonsue/labchat/model/menu"
)

// MenuRepository struct definition
type MenuRepository struct{}

// NewMenuRepository does several services according to InMemoryDB
func NewMenuRepository() menu.Repository {
	return &MenuRepository{}
}
