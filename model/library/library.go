package library

type LoginInfo struct {
	UserKey    string // Kakao Userkey
	LoginId    string // Library login ID
	Password   string // Library login PW
	LoginToken string // Library login token
}

type Repository interface {
	Find(userkey string) (*LoginInfo, error)
	Store(key *LoginInfo) error
	StoreToken(id string, token string) error
	Clean() error
}
