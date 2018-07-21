package inmem

import (
	"github.com/yoonsue/labchat/model/menu"
)

type MenuRepository struct{}

func NewMenuRepository() menu.Repository {
	return &MenuRepository{}
}
