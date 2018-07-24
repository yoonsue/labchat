package mongo

import (
	"github.com/yoonsue/labchat/model/menu"
)

// MenuRepository
type MenuRepository struct{}

// NewMenuRepository
func NewMenuRepository() menu.Repository {
	return &MenuRepository{}
}
