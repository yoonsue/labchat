package library

// LoginInfo is a struct for storing the user login information.
type LoginInfo struct {
	UserKey  string // Kakao Userkey
	LoginID  string // Library login ID
	Password string // Library login PW
	// LoginToken string // Library login token
	JSessionID string // Library jsession token
}

// Repository interface defines functions which must be implemented in the library repository.
type Repository interface {
	Find(userkey string) (*LoginInfo, error)
	Store(key *LoginInfo) error
	StoreJSessionID(loginInfo *LoginInfo, JSession string) error
	Clean() error
}
