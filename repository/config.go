package repository

// DBConfig is the value holder of the configurations of labchat database.
type DBConfig struct {
	// TODO: implementation.
	env string
}

// DefaultConfig returns the default setting of the labchat configuration.
func DefaultConfig() *DBConfig {
	return &DBConfig{
		env: "none",
	}
}
