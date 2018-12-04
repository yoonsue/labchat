package library

type LoginInfo struct {
	UserKey  string // Kakao Userkey
	LoginID  string // Library login ID
	Password string // Library login PW
	// LoginToken string // Library login token
	JSessionID string // Library jsession token
}

type Repository interface {
	Find(userkey string) (*LoginInfo, error)
	Store(key *LoginInfo) error
	StoreJSessionID(loginInfo *LoginInfo, JSession string) error
	Clean() error
}
