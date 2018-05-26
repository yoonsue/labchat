package server

import (
	"fmt"
	"net/http"
	"os"
)

// Address is for server address.
type Address string

// Server provides http service for the labchat service.
type Server struct {
	// TODO: implementation.
	cfg *Config
}

// NewServer creates a new labchat server with the given configuration.
func NewServer(cfg *Config) (srv *Server, err error) {
	// TODO: implementation.
	return &Server{
		cfg: cfg,
	}, nil
}

// Start runs the server and starts handling for incoming requests. All the
// long-running server functionality should be implemented in goroutines.
func (s *Server) Start() {
	// TODO: implementation.
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	// TODO: need to halt goroutine when the program is stopped.
	go http.ListenAndServe(s.cfg.Address, nil)
}

func loadMsg(filepath string) bool {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		fmt.Print("Error: openFile %s\n", err)
		return false
	}
	defer file.Close()

	n := 0
	s := "Hello"
	n, err = file.Write([]byte(s))
	if err != nil {
		fmt.Print("Error: write %s\n", err)
		return false
	}
	fmt.Println(n, " byte saved in ", filepath)

	fi, err := file.Stat()
	if err != nil {
		fmt.Print("Error: file stat %s\n", err)
		return false
	}

	var data = make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET)

	n, err = file.Read(data)
	if err != nil {
		fmt.Print("Error: read %s\n", err)
		return false
	}

	fmt.Println(n, " byte read from ", filepath)
	fmt.Println(string(data))
	return true
}
