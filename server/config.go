package server

// Config is the value holder of the configurations of labchat server.
type Config struct {
	// TODO: implementation.
	Address string
}

// DefaultConfig returns the default setting of the labchat configuration.
func DefaultConfig() *Config {
	return &Config{
		Address: "localhost:4748",
	}
}
