package birthday

// Brithday = VO
type Birthday struct {
	Name     string
	Birthday int
}

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key string) (*Birthday, error)
	Store(key *Birthday) error
}
