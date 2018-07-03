package server

import (
	"encoding/json"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/yoonsue/labchat/function"

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
	http.HandleFunc("/labchat/", handlehttp)

	// TODO: need to halt goroutine when the program is stopped.
	go http.ListenAndServe(s.cfg.Address, nil)
}

//// Below stuctures are created following 'Kakao Api specifications'
//// more information at 'https://github.com/plusfriend/auto_reply'

// 'keyboard' contains information about buttons that are in the field keyboard
// GET	http://your_server_url/keyboard
type keyboard struct {
	Type    string   `json:"type"`
	Buttons []string `json:"buttons"`
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
		// CASE1 Type: 'text'
		initkeyboard := keyboard{Type: "text"}

		// // CASE2 Type: 'buttons'
		// initkeyboard := make([]keyboard, 1)
		// initkeyboard[0].Type = "buttons"
		// initkeyboard[0].Buttons = []string{"option1", "option2", "option3"}

		resp, err := json.Marshal(initkeyboard)
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

		// check mesKey exist
		msgCon := messageKey(msg.Content)
		// if there is no msgKey, return same msg
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

	// friend information
	// curl -XPOST 'https://:your_server_url/friend' -d '{"user_key" : "HASHED_USER_KEY" }'
	if r.Method == "POST" && r.URL.Path == "/labchat/friend" {
		body, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Println(errors.Wrap(err, "failed to read body of /friend"))
		}
		log.Printf("body: %s\n", string(body))

		var usr user
		if err := json.Unmarshal(body, &usr); err != nil {
			log.Println(errors.Wrap(err, "failed to unmarshal body of /friend"))
		}
		log.Printf("Friend %s joined\n", usr.UserKey)
		return
	}

	if r.Method == "DELETE" {
		split := strings.Split(r.URL.Path, ":")

		// chatroom delete by admin
		// curl -XDELETE 'https://:your_server_url/chat_room/HASHED_USER_KEY'
		if split[0] == "/labchat/friend/" {
			log.Printf("user %s deleted", split[1])
		}

		// chatroom deleted by user
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
	"name":  "LABchat",
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
	if tokens[0] == "lab" || tokens[0] == "Lab" || tokens[0] == "LAB" {
		if len(tokens) < 2 {
			return "no command"
		}
		// TODO
		if tokens[1] == "status" {
			status := function.ServerCheck()
			time := status.TimeStamp
			strtime := time.Format("2006-01-02 15:04:05 ")
			temp := status.Temperature
			return ("TIME : " + strtime + "\n TEMP : " + strconv.FormatFloat(float64(temp), 'g', -1, 64) + " C")
		}
		if tokens[1] == "menu" {
			menu := function.GetMenu("http://www.hanyang.ac.kr/web/www/-255")
			menuS := string(menu)
			return ("cafeteria menu : " + menuS)
		}
	}
	return strings.Join(tokens, " ") + "....????"
}
