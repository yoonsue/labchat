package server

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
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
	filepath := "msg.json"
	if !loadJson(filepath) {
		log.Println("Error: func loadMsg failed")
	}
	log.Println("Success: loadMsg")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	// TODO: need to halt goroutine when the program is stopped.
	go http.ListenAndServe(s.cfg.Address, nil)
}

//// Below stuctures are created following 'Kakao Api specifications'
//// more information at 'https://github.com/plusfriend/auto_reply'

// 'keyboard' contains information about buttons that are in the field keyboard
// GET	http://your_server_url/keyboard
type keyboard struct {
	Type string `json:"type"`
}

// 'message'
// POST	http://your_server_url/message
type message struct {
	UserKey string `json:"user_key"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// response struct meaning
type resptext struct {
	Text string `json:"text"`
}
type response struct {
	Message resptext `json:"message"`
}
type user struct {
	UserKey string `json:"user_key"`
}

// POST	http://your_server_url/friend
// DELETE	http://your_server_url/friend/:user_key

// DELETE	http://:your_server_url/chat_room/:user_key

func handlehttp(w http.ResponseWriter, r *http.Request) {
	log.Printf("received: %s\t %s\n", r.Method, html.EscapeString(r.URL.Path))

	if r.Method == "GET" && r.URL.Path == "/labchat/keyboard" {
		resp, err := json.Marshal(keyboard{Type: "text"})
		if err != nil {
			log.Println(errors.Wrap(err, "failed to marshal 'keyboard'"))
		}
		fmt.Fprint(w, string(resp))
		return
	}

	if r.Method == "POST" && r.URL.Path == "/labchat/message" {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Println(errors.Wrap(err, "failed to read body of /message"))
		}
		log.Printf("body: %s\n", string(body))

		var msg message
		if err := json.Unmarshal(body, &msg); err != nil {
			log.Println(errors.Wrap(err, "failed to unmarshal /message"))
		}
		return
	}

	if r.Method == "POST" && r.URL.Path == "/labchat/friend" {
		return
	}

	if r.Method == "DELETE" {
		split := strings.Split(r.URL.Path, ":")
		if split[0] == "/labchat/friend/" {
			log.Printf("user %s deleted", split[1])
		}
		if split[0] == "/labchat/chat_room/" {
			log.Printf("user %s leaved", split[2])
		}
		return
	}
}
