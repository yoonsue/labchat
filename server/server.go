package server

import (
	"encoding/json"
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
	// filepath := "msg.json"
	// if !loadJson(filepath) {
	// 	log.Println("Error: func loadMsg failed")
	// }
	// log.Println("Success: loadMsg")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/labchat/", handlehttp)

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

func handlehttp(w http.ResponseWriter, r *http.Request) {
	log.Printf("received: %s\t %s\n", r.Method, html.EscapeString(r.URL.Path))

	// curl -XGET 'https://:your_server_url/keyboard'
	if r.Method == "GET" && r.URL.Path == "/labchat/keyboard" {
		resp, err := json.Marshal(keyboard{Type: "text"})
		if err != nil {
			log.Println(errors.Wrap(err, "failed to marshal 'keyboard'"))
		}
		w.Write([]byte(string(resp) + "\n"))
		return
	}

	//curl -XPOST 'https://:your_server_url/message' -d '{
	//   "user_key": "encryptedUserKey",
	//   "type": "text",
	//   "content": "차량번호등록"
	// }'
	if r.Method == "POST" && r.URL.Path == "/labchat/message" {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Println(errors.Wrap(err, "failed to read body of /message"))
		}
		log.Printf("body: %s\n", string(body))

		// input unmarshaled body at message
		var msg message
		if err := json.Unmarshal(body, &msg); err != nil {
			log.Println(errors.Wrap(err, "failed to unmarshal /message"))
		}

		// check meskey exist
		msgCon := messageKey(msg.Content)
		// if there is no msg key, return same msg
		if msgCon == "none" {
			msgCon = msg.Content
		}

		// response
		// {
		// 	"message":{
		// 		"text" : "귀하의 차량이 성공적으로 등록되었습니다. 축하합니다!"
		// 	}
		// }
		remsg := msgFor(strings.Fields(msgCon))
		resp, err := json.Marshal(response{
			Message: resptext{
				Text: remsg}})
		if err != nil {
			log.Println(errors.Wrap(err, "failed to marshal response"))
		}
		log.Printf("send %s\n", string(resp))
		w.Write([]byte(string(resp) + "\n"))
		return
	}

	// curl -XPOST 'https://:your_server_url/friend' -d '{"user_key" : "HASHED_USER_KEY" }'
	if r.Method == "POST" && r.URL.Path == "/labchat/friend" {
		return
	}

	if r.Method == "DELETE" {
		split := strings.Split(r.URL.Path, ":")

		// curl -XDELETE 'https://:your_server_url/chat_room/HASHED_USER_KEY'
		if split[0] == "/labchat/friend/" {
			log.Printf("user %s deleted", split[1])
		}

		// DELETE	http://:your_server_url/chat_room/:user_key
		if split[0] == "/labchat/chat_room/" {
			log.Printf("user %s leaved", split[2])
		}
		return
	}
}

var messageKeyMap = map[string]string{
	"hi":    "hello",
	"hello": "hi",
}

func messageKey(rawmessage string) string {
	key, result := messageKeyMap[rawmessage]
	if !result {
		return "none"
	}
	return key
}

func msgFor(tokens []string) string {
	// exec command
	if tokens[0] == "ex" {
		if len(tokens) < 2 {
			return "no command"
		}
		// TODO
	}
	return strings.Join(tokens, " ") + "....????"
}
