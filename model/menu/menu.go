package menu

// Menu = VO
type Menu struct {
	Restaurant Restaurant
	TodayMenu  TodayMenu
}

type Restaurant string
type TodayMenu string

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key Restaurant) (TodayMenu, error)
	Store(key Menu) error
}
