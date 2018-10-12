package birthday

// Birthday = VO
type Birthday struct {
	Name     string
	Birthday string
	Age      int
}

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key string) (*Birthday, error)
	Store(key *Birthday) error
}
