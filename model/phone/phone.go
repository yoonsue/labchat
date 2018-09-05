package phone

// Phone = VO
type Phone struct {
	Department Department
	Extension  Extension
}

// Department ...
type Department string

// Extension ...
type Extension int

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key Department) (*Phone, error)
	Store(key *Phone) error
}
