package location

// Location = VO
type Location struct {
	Name     string
	Location string
}

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key string) ([]*Location, error)
	Store(key *Location) error
	Clean() error
}
