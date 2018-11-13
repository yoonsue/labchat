package library

type LoginInfo struct {
	LoginId  string
	Password string
}

type Repository interface {
	Find(id string, pw string) (*LoginInfo, error)
	Store(key *LoginInfo) error
	Clean() error
}
