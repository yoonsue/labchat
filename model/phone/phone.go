package phone

// Phone = VO
type Phone struct {
	Department string
	Extension  int
}

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key string) (*Phone, error)
	Store(key *Phone) error
}
