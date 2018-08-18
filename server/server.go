package server

import (
	"encoding/json"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/function/menu"
	"github.com/yoonsue/labchat/function/status"
)

// Address is for server address.
type Address string

// Server provides http service for the labchat service.
type Server struct {
	// TODO: implementation.
	cfg         *Config
	menuService menu.Service
}

// NewServer creates a new labchat server with the given configuration.
func NewServer(cfg *Config, ms menu.Service) (srv *Server, err error) {
	// TODO: implementation.
	return &Server{
		cfg:         cfg,
		menuService: ms,
	}, nil
}

// Start runs the server and starts handling for incoming requests. All the
// long-running server functionality should be implemented in goroutines.
func (s *Server) Start() {
	// TODO: implementation.
	http.HandleFunc("/", s.handleHTTP)

	// TODO: need to halt goroutine when the program is stopped.
	go http.ListenAndServe(s.cfg.Address, nil)
}

//// Below stuctures are created following 'Kakao API specifications'
//// more information at 'https://github.com/plusfriend/auto_reply'

// keyboard has type that two initial options (button or text)
type keyboard struct {
	Type    string   `json:"type"`
	Buttons []string `json:"buttons"`
}

// message contains information about UserKey, Type and Content
type message struct {
	UserKey string `json:"user_key"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// respText is used when response type is text
type respText struct {
	Text string `json:"text"`
}

// response contains Message for respText
type response struct {
	Message respText `json:"message"`
}

// user contains UserKey for user_key of message
type user struct {
	UserKey string `json:"user_key"`
}

// handleHTTP is requested handler of Kakao API (RESTful API)
func (s *Server) handleHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("received: %s\t %s\n", r.Method, html.EscapeString(r.URL.Path))

	rou := mux.NewRouter()

	rou.HandleFunc("/", s.test)

	rou.HandleFunc("/keyboard", s.keyboardInit).Methods("GET")
	rou.HandleFunc("/message", s.messageSet).Methods("POST")
	rou.HandleFunc("/friend", s.newFriend).Methods("POST")
	rou.HandleFunc("/friend/", s.removeFriend).Methods("DELETE")
	rou.HandleFunc("/chat_room/", s.leaveChatRoom).Methods("DELETE")
	err := rou.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			log.Println(errors.Wrap(err, "Path regexp:"), pathTemplate)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			log.Println(errors.Wrap(err, "Queries templates: "), queriesTemplates)
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			log.Println(errors.Wrap(err, "Queries regexps: "), queriesRegexps)
		}
		methods, err := route.GetMethods()
		if err == nil {
			log.Println(errors.Wrap(err, "Methods: "), methods)
		}
		log.Println()
		return nil
	})
	if err != nil {
		log.Println(errors.Wrap(err, "ERROR: "))
	}

	// http.Handle("/", rou)
	// return rou
}

func (s *Server) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TESTING...\n"))
	return
}

// curl -XGET 'https://:your_server_url/keyboard'
func (s *Server) keyboardInit(w http.ResponseWriter, r *http.Request) {
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

// curl -XPOST 'https://:your_server_url/message' -d '{
//   "user_key": "encryptedUserKey",
//   "type": "text",
//   "content": "차량번호등록"
// }'
func (s *Server) messageSet(w http.ResponseWriter, r *http.Request) {
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

	msgCon := messageKey(msg.Content)
	// if there is no msgCon, return same msg
	if msgCon == "none" {
		msgCon = msg.Content
	}

	remsg := s.msgFor(msgCon)
	resp, err := json.Marshal(response{
		Message: respText{
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
func (s *Server) newFriend(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) removeFriend(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, ":")

	// chatroom delete by admin
	// curl -XDELETE 'https://:your_server_url/chat_room/HASHED_USER_KEY'
	if split[0] == "/friend/" {
		log.Printf("user %s deleted", split[1])
	}
}

func (s *Server) leaveChatRoom(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, ":")

	// chatroom deleted by user
	// DELETE	http://:your_server_url/chat_room/:user_key
	if split[0] == "/chat_room/" {
		log.Printf("user %s leaved", split[2])
	}
	return
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

// menuURLMap
// var menuURLMap = map[string]string{"교직원식당": "http://www.hanyang.ac.kr/web/www/-254", "학생식당": "http://www.hanyang.ac.kr/web/www/-255", "창업보육센터": "http://www.hanyang.ac.kr/web/www/-258", "창의인재원식당": "http://www.hanyang.ac.kr/web/www/-256"}
var menuURLMap = []string{"http://www.hanyang.ac.kr/web/www/-254", "http://www.hanyang.ac.kr/web/www/-255", "http://www.hanyang.ac.kr/web/www/-258", "http://www.hanyang.ac.kr/web/www/-256"}

func (s *Server) msgFor(request string) string {
	// var s *Server
	if request == "status" {
		c := status.ServerCheck()

		str := "TIME : " + c.Time()
		str = str + "\n"
		str = str + "TEMP : " + c.Temperature.String()

		return str
	}
	if request == "menu" {
		str := ""
		for _, menuURL := range menuURLMap {
			menuRest := s.menuService.GetSchool(menuURL)
			str += string(menuRest.Restaurant)
			str += "\n"
			str += string(menuRest.TodayMenu)
			str += "\n\n"
		}
		return str
	}
	return request + "....????"
}
