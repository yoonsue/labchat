package menu

// Menu = VO
type Menu struct {
	Restaurant Restaurant
	TodayMenu  TodayMenu
}

// Restaurant ...
type Restaurant string

// TodayMenu ...
type TodayMenu string

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key Restaurant) (*Menu, error)
	Store(key *Menu) error
}
