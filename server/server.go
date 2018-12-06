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
	"github.com/yoonsue/labchat/function/library"
	"github.com/yoonsue/labchat/function/location"
	"github.com/yoonsue/labchat/function/menu"
	"github.com/yoonsue/labchat/function/phone"
	"github.com/yoonsue/labchat/function/status"
	statusModel "github.com/yoonsue/labchat/model/status"
)

// Address is for server address.
type Address string

// Server provides http service for the labchat service.
type Server struct {
	// TODO: implementation.
	currentTime     string
	cfg             *Config
	menuService     menu.Service
	phoneService    phone.Service
	statusService   status.Service
	birthdayService birthday.Service
	locationService location.Service
	libraryService  library.Service
}

// NewServer creates a new labchat server with the given configuration.
func NewServer(curTime string, cfg *Config, ms menu.Service, ps phone.Service, ss status.Service, bs birthday.Service, ls location.Service, libs library.Service) (srv *Server, err error) {
	return &Server{
		currentTime:     curTime,
		cfg:             cfg,
		menuService:     ms,
		phoneService:    ps,
		statusService:   ss,
		birthdayService: bs,
		locationService: ls,
		libraryService:  libs,
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
	w.Write([]byte("Hello, my friend"))
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

	switch request[0] {

	case "도움말":
		return s.help()

	case "시작하기":
		return s.start()

	case "status", "Status", "상태", "서버":
		return s.status()

	case "menu", "Menu", "메뉴", "학식":
		return s.menu()

	case "phone", "Phone", "내선", "번호", "내선번호":
		return s.phone(request)

	case "birthday", "Birthday", "생일":
		return s.birthday(request)

	case "location", "위치", "주소":
		return s.location(request)

	case "library", "도서관", "도서", "연장":
		return s.library(request)
	}

	str := strings.Join(request, " ") + "....????"
	return str
}

func (s *Server) help() string {
	str := ""
	todayIsBirthdayList := s.birthdayService.CheckBirthday(s.currentTime)
	for _, tmp := range todayIsBirthdayList {
		str += "(축하)오늘은 "
		str += tmp.Name
		str += "의 생일입니다.(축하)\n"
	}

	str += "LABchat에 오신걸 환영합니다.\nLABchat은 다음과 같은 기능을 제공합니다.\n - 서버상태 정보(status)\n - 내선번호 검색 기능(phone 사이버피지컬)\n - 컴퓨터공학과 교수진 및 연구실 주소 검색 기능(location 사이버피지컬)\n - 교내식당 메뉴 정보 제공(menu)\n - 연구실 내 구성원 생일 정보 제공(birthday 이름)\n"
	return str
}

func (s *Server) start() string {
	str := ""
	todayIsBirthdayList := s.birthdayService.CheckBirthday(s.currentTime)
	for _, tmp := range todayIsBirthdayList {
		str += "(축하)오늘은 "
		str += tmp.Name
		str += "의 생일입니다.(축하)\n"
	}

	str += "필요한 기능을 사용해보세요."
	return str
}

func (s *Server) status() string {
	str := ""
	c := s.statusService.ServerCheck()
	// str := "TIME : " + c.Time()
	str += "TEMP : " + c.Temperature.String()
	str += "\nUPTIME: " + statusModel.FmtDuration(c.Uptime)
	return str
}

func (s *Server) menu() string {
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

func (s *Server) phone(request []string) string {
	str := ""
	if len(request) < 2 {
		return "no department"
	}

	p, _ := s.phoneService.GetPhone(request[1])
	if p == nil {
		str += "No result from the given department"
	} else {
		for _, val := range p {
			str += string(val.Department)
			str += "\t"
			str += val.Extension
			str += "\n"
		}
	}
	return str
}

func (s *Server) birthday(request []string) string {
	str := ""
	if len(request) < 2 {
		return "no name"
	}
	name := request[1]
	// CHECKPOINT: finding all birthday is necessary?
	// if name == "모두" || name == "전원" || name == "연구실" {
	// 	b. _ := s.birthdayService.GetAllBirthday()
	// }
	b, _ := s.birthdayService.GetBirthday(name)
	if b == nil {
		str += "No result from the given name"
	} else {
		str += b.Name
		str += "\t나이: "
		str += strconv.Itoa(b.GetAge())
		str += "\t생일: "
		str += b.GetBirth()
	}
	return str
}

func (s *Server) location(request []string) string {
	str := ""
	if len(request) < 2 {
		return "no name"
	}
	name := request[1]
	l, _ := s.locationService.GetLocation(name)
	if l == nil {
		str += "No result from the given location"
	} else {
		for _, val := range l {
			str += val.Name
			str += " 위치: "
			str += val.Location
			str += "\n"
		}
	}
	return str
}

func (s *Server) library(request []string) string {
	str := ""
	if len(request) < 3 {
		return "no id and pw"
	}
	id := request[1]
	pw := request[2]
	l, _ := s.libraryService.Login(id, pw)
	if l == nil {
		str += "No result from the given id and pw"
	} else {
		str += l.LoginID
		str += "님의 도서 "
		str += "3" /////////
		str += " 권("
		str += "도서명1, 도서명2, 도서명3" /////////
		str += "을 자동연장하였습니다."
		str += "\nJSessionID: " + l.JSessionID
	}
	return str
}
