package server

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	// TODO: implementation.
	// how to collarborate the fixed default address with config.go
	config := &Config{
		Address: "localhost:4748",
	}

	cfg := DefaultConfig()

	if cfg.Address != config.Address {
		t.Errorf("expected %s, got %s", config.Address, cfg.Address)
	}
}
