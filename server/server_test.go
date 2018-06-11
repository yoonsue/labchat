package server

import (
	"net/http"
	"testing"

	"github.com/yoonsue/labchat/server"
)

func TestServer(t *testing.T) {
	// TODO: implementation.
	serverConfig := server.DefaultConfig()
	testserver, err := &Server{
		cfg: serverConfig,
	}, nil

	testserver.Start()

	resp, err := http.Get("http://" + testserver + "/labchat/keyboard")
	if err != nil {
		// handle err
		t.Fatal(err)
	}
	defer resp.Body.Close()
}
