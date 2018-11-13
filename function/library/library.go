package library

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	libraryModel "github.com/yoonsue/labchat/model/library"
)

type Address string

// Proxy provides http service for the library service.
type Proxy struct {
	// TODO: implementation.
	currentTime string
	cfg         Address
}

func (p *Proxy) Start() *mux.Router {
	rou := mux.NewRouter()

	rou.HandleFunc("/api/login", p.loginHandler).Methods("POST")
}

type message struct {
	UserKey string `json:"user_key"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type loginRequest struct {
	loginId  string `json:"loginId"`
	password string `json:"password"`
}

// response contains Message for respText
type response struct {
	Message respLogin `json:"message"`
}

// respLogin is used when response type is text
type respLogin struct {
	AccessToken string `json:"accessToken"`
}

// Request
// curl -H 'Content-Type: application/json;charset=UTF-8'
// 	-XPOST 'http://lib.hanyang.ac.kr/pyxis-api/api/login'
// 	-d '{"loginId": "----",  "password": "----"}'
func (p *Proxy) loginHandler(w http.ResponseWriter, r *http.Request) string {
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

	var Token string
	initLogin := libraryModel.LoginInfo{}

	aT := strings.Fields(msg.Content)
	resp, err := json.Marshal(response{
		Message: respLogin{
			AccessToken: aT}})
	if err != nil {
		log.Println(errors.Wrap(err, "failed to marshal response"))
	}

	return Token
}

// Response
// {
// 	"success":true,
// 	"code":"success.loggedIn",
// 	"message":"로그인되었습니다.",
// 	"data":{
// 		"isPortalLogin":true,
// 		"alternativeId":"00000000-0000-0000-0000-000000000000",
// 		"accessToken":"00000000000000000000000000000000", 			// KEY ELEMENT 32bit
// 		"parentDept":{
// 			"id":65,
// 			"code":"H0000476",
// 			"name":"대학원"
// 		},"lastUpdated":"2018-08-23 10:40:50",
// 		"patronState":{
// 			"id":8,
// 			"name":"재학"
// 		},"isPrivacyPolicyAgree":true,
// 		"printMemberNo":"0000000000", 						// 학번
// 		"id":000000,
// 		"isExpired":false,
// 		"disableServices":["WORKER_RECALL","LECTURE","USER_INFO"],
// 		"hasFamily":false,
// 		"name":"000",
// 		"branch":{
// 			"id":9,
// 			"name":"ERICA학술정보관",
// 			"alias":"ERICA",
// 			"libraryCode":"241050"
// 		},"multiTypePatrons":[],
// 		"availableHomepages":[1,2,3,4,5,6,7],
// 		"dept":{
// 			"id":653,
// 			"code":"H0000706",
// 			"name":"컴퓨터공학과"
// 		},"patronType":{"id":2,"name":"대학원"},
// 		"memberNo":"0000000000"								// 학번
// 	}
// }
