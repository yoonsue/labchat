package phone

// Phone = VO
type Phone struct {
	Department Department
	Extension  string
}

// Department ...
type Department string

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key Department) ([]*Phone, error)
	Store(key *Phone) error
	Clean() error
}

// ToString converts department type to string.
func (d Department) ToString() string {
	return string(d)
}
