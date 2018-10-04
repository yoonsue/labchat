package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/yoonsue/labchat/function/birthday"
	"github.com/yoonsue/labchat/function/menu"
	"github.com/yoonsue/labchat/function/phone"
	"github.com/yoonsue/labchat/function/status"
	phoneModel "github.com/yoonsue/labchat/model/phone"
	statusModel "github.com/yoonsue/labchat/model/status"
)

// Address is for server address.
type Address string

// Server provides http service for the labchat service.
type Server struct {
	// TODO: implementation.
	cfg             *Config
	menuService     menu.Service
	phoneService    phone.Service
	statusService   status.Service
	birthdayService birthday.Service
}

// NewServer creates a new labchat server with the given configuration.
func NewServer(cfg *Config, ms menu.Service, ps phone.Service, ss status.Service, bs birthday.Service) (srv *Server, err error) {
	return &Server{
		cfg:             cfg,
		menuService:     ms,
		phoneService:    ps,
		statusService:   ss,
		birthdayService: bs,
	}, nil
}

// Start runs the server and starts handling for incoming requests. All the
// long-running server functionality should be implemented in goroutines.
func (s *Server) Start() *mux.Router {
	// TODO: implementation.
	rou := mux.NewRouter()

	rou.HandleFunc("/keyboard", s.keyboardHandler).Methods("GET")
	rou.HandleFunc("/message", s.messageHandler).Methods("POST")
	rou.HandleFunc("/friend", s.friendHandler).Methods("POST")
	rou.HandleFunc("/friend/{id}", s.friendDeleteHandler).Methods("DELETE")
	rou.HandleFunc("/chat_room/{id}", s.chatroomDeleteHandler).Methods("DELETE")

	// TODO: need to halt goroutine when the program is stopped.
	go http.ListenAndServe(s.cfg.Address, rou)
	return rou
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

// respButton is used when response type is text
type respButton struct {
	Button string `json:"buttons"`
}

// response contains Message for respText
type response struct {
	Message respText `json:"message"`
}

// user contains UserKey for user_key of message
type user struct {
	UserKey string `json:"user_key"`
}

// curl -XGET 'https://:your_server_url/keyboard'
func (s *Server) keyboardHandler(w http.ResponseWriter, r *http.Request) {
	// Type: 'text' or 'buttons' - default is 'text'.
	initkeyboard := keyboard{
		Type:    "buttons",
		Buttons: []string{"도움말", "시작하기"},
	}

	resp, err := json.Marshal(initkeyboard)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to marshal 'keyboard'"))
	}
	w.Write([]byte(string(resp) + "\n"))
	return
}

// curl -XPOST 'https://:your_server_url/message' -d '{  "user_key": "encryptedUserKey",  "type": "text",  "content": "차량번호등록"}'
func (s *Server) messageHandler(w http.ResponseWriter, r *http.Request) {
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

	remsg := s.msgFor(strings.Fields(msgCon))
	resp, err := json.Marshal(response{
		Message: respText{
			Text: remsg}})
	if err != nil {
		log.Println(errors.Wrap(err, "failed to marshal response"))
	}
	log.Printf("send %s\n", string(resp))
	w.Write([]byte(string(resp)))
	return
}

// friend information
// curl -XPOST 'https://:your_server_url/friend' -d '{"user_key" : "HASHED_USER_KEY" }'
func (s *Server) friendHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(string("Hello, my friend")))
	log.Printf("Friend %s joined\n", usr.UserKey)
	return
}

// friend delete by admin
// curl -XDELETE 'https://:your_server_url/friend/:user_key'
func (s *Server) friendDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userKey := vars["id"]
	log.Printf("user %s deleted", userKey)
	w.Write([]byte(string("\nuser deleted")))
	return
}

// chatroom deleted by user
// curl -XDELETE 'https://:your_server_url/chat_room/HASHED_USER_KEY'
func (s *Server) chatroomDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userKey := vars["id"]
	log.Printf("user %s leaved", userKey)
	w.Write([]byte(string("\nchatroom deleted")))
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
//// var menuURLMap = map[string]string{"교직원식당": "http://www.hanyang.ac.kr/web/www/re11", "학생식당": "http://www.hanyang.ac.kr/web/www/re12", "창업보육센터": "http://www.hanyang.ac.kr/web/www/re13", "창의인재원식당": "http://www.hanyang.ac.kr/web/www/re15"}
var menuURLMap = []string{"https://www.hanyang.ac.kr/web/www/re11",
	"http://www.hanyang.ac.kr/web/www/re12",
	"http://www.hanyang.ac.kr/web/www/re13",
	"http://www.hanyang.ac.kr/web/www/re15"}

func (s *Server) msgFor(request []string) string {
	if request[0] == "도움말" {
		str := "LABchat에 오신걸 환영합니다.\nLABchat은 다음과 같은 기능을 제공합니다.\n - 서버상태 정보(status)\n - 내선번호 검색 기능(phone 사이버피지컬시스템연구실(CPS))\n - 교내식당 메뉴 정보 제공(menu)\n"
		return str
	}
	if request[0] == "시작하기" {
		str := "필요한 기능을 사용해보세요."
		return str
	}
	if request[0] == "status" || request[0] == "Status" || request[0] == "상태" || request[0] == "서버" {

		c := s.statusService.ServerCheck()

		// str := "TIME : " + c.Time()
		str := "TEMP : " + c.Temperature.String()
		str += "\nUPTIME: " + statusModel.FmtDuration(c.Uptime)

		return str
	}
	if request[0] == "menu" || request[0] == "Menu" || request[0] == "학식" || request[0] == "메뉴" {
		str := ""
		for _, menuURL := range menuURLMap {
			menu := s.menuService.GetSchool(menuURL)
			str += string(menu.Restaurant)
			str += "\n"
			str += string(menu.TodayMenu)
			str += "\n\n"
		}
		return str
	}
	if request[0] == "phone" || request[0] == "Phone" || request[0] == "내선" || request[0] == "번호" || request[0] == "내선번호" {
		if len(request) < 2 {
			return "no department"
		}
		department := phoneModel.Department(request[1])
		p, _ := s.phoneService.GetPhone(department)
		str := ""
		if p == nil {
			str += "No result.."
		} else {
			str += string(p.Department)
			str += "\t"
			str += p.Extension
		}
		return str
	}
	if request[0] == "birthday" || request[0] == "Birthday" || request[0] == "생일" {
		if len(request) < 2 {
			return "no name"
		}
		name := request[1]
		b, _ := s.birthdayService.GetBirthday(name)
		str := ""
		if b == nil {
			str += "No result.."
		} else {
			str += b.Name
			str += "\t"
			str += strconv.Itoa(b.Birthday)
		}
		return str
	}
	return strings.Join(request, " ") + "....????"
}
