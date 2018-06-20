package function

// Targetserver is the value holder of the configurations of target server.
type Targetserver struct {
	// TODO: implementation.
	Address     string
	Temperature float64
}

// DefaultServer returns the default setting of the target server.
func DefaultServer() *Targetserver {
	return &Targetserver{
		Address:     "none",
		Temperature: -1,
	}
}
