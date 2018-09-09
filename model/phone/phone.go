package phone

// Phone = VO
type Phone struct {
	Department Department
	Extension  int
}

// Department ...
type Department string

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key Department) (*Phone, error)
	Store(key *Phone) error
}
