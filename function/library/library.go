package library

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model/library"
)

const defaultLibraryAddress = "https://lib.hanyang.ac.kr/pyxis-api"

type Service interface {
	// TO BE IMPLEMENTED: it would be a kakaotalk api userkey

	Login(id string, pw string) (*library.LoginInfo, error)
	// GetLoginInfo(string, string) (*library.LoginInfo, error)
	// Login(string) error
}

func NewService(r library.Repository, fpath string) Service {
	s := &service{
		libraryLoginList: r,
	}
	s.intialStore(fpath)
	return s
}

type service struct {
	libraryLoginList library.Repository
}

// func (s *service) Login(id string, pw string) (*library.LoginInfo, error) { //return accessToken
// 	p, err := NewProxy(time.Now().Local().Format("2006-01-02"), Address(defaultLibraryAddress), s)
// 	if err != nil {
// 		log.Println(errors.Wrap(err, "failed to start new proxy"))
// 	}
// 	p.Start()
// 	return nil, err
// }

func (s *service) GetDueDate(userkey string) (string, error) {
	userLoginInfo, err := s.libraryLoginList.Find(userkey)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get login information"))
		return "", err
	}
	return userLoginInfo.LoginToken, nil
}

type Address string

// Proxy provides http service for the library service.
type Proxy struct {
	// TODO: implementation.
	currentTime    string
	cfg            Address
	libraryService Service
}

func NewProxy(curTime string, cfg Address, ls Service) (p *Proxy, err error) {
	return &Proxy{
		currentTime:    curTime,
		cfg:            cfg,
		libraryService: ls,
	}, nil
}

func (p *Proxy) Start() *mux.Router {
	rou := mux.NewRouter()

	// rou.HandleFunc("/api/login", p.loginHandler).Methods("POST")

	return rou
}

type message struct {
	UserKey string `json:"user_key"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type loginRequest struct {
	loginId  string `json:"loginId"`
	password string `json:"password"`

	/// TO BE COMFIRMED
	// jsessionID     string `json:"JSESSIONID"`
	// pyxisAuthToken string `json:"pyxis-auth-token"`
}

// response contains Message for respText
type response struct {
	success bool   `json:"success"`
	code    string `json:"code"`
	message bool   `json:"message"`
	data    data   `json:"data"`
}

type data struct {
	isPortalLogin        bool        `json:"isPortalLogin"`
	alternativeId        string      `json:"alternativeId"`
	accessToken          string      `json:"accessToken"`
	parentDept           parentDept  `json:"parentDept"`
	lastUpdated          string      `json:"lastUpdated"`
	patronState          patronState `json:"patronState"`
	isPrivacyPolicyAgree bool        `json:"isPrivacyPolicyAgree"` //:true,
	printMemberNo        int         `json:"printMemberNo"`        //:"0000000000", 						// 학번
	id                   string      `json:"id"`                   //:000000,
	isExpired            bool        `json:"isExpired"`            //:false,
	disableServices      []string    `json:"disableServices"`      //:["WORKER_RECALL","LECTURE","USER_INFO"],
	hasFamily            bool        `json:"hasFamily"`            //:false,
	name                 string      `json:"name"`                 //:"000",
	branch               branch      `json:"branch"`
	multiTypePatrons     []string    `json:"multiTypePatrons"`   //:[],
	availableHomepages   []int       `json:"availableHomepages"` //:[1,2,3,4,5,6,7],
	dept                 dept        `json:"dept"`
	patronType           patronType  `json:"patronType"`
	memberNo             int         `json:"memberNo"` //: 학번
}

// {"success":true,
// "code":"success.loggedIn",
// "message":"로그인되었습니다.",
// "data":{
// 	"isPortalLogin":true,
// 	"alternativeId":"76f26cac-82db-40d5-9294-f111590efaaa",
// 	"accessToken":"kj63vhmk9gsvr6pmbofb6o6kmapsjdjc",
// 	"parentDept":{"id":65,"code":"H0000476","name":"대학원"},
// 	"lastUpdated":"2018-08-23 10:40:50",
// 	"patronState":{"id":8,"name":"재학"},
// 	"isPrivacyPolicyAgree":true,
// 	"printMemberNo":"20181217951",
// 	"id":413295,
// 	"isExpired":false,
// 	"disableServices":["WORKER_RECALL","LECTURE","USER_INFO"],
// 	"hasFamily":false,
// 	"name":"조윤수",
// 	"branch":{
// 		"id":9,"name":"ERICA학술정보관","alias":"ERICA",
// 		"libraryCode":"241050"},"multiTypePatrons":[],
// 		"availableHomepages":[1,2,3,4,5,6,7],
// 		"dept":{"id":653,"code":"H0000706","name":"컴퓨터공학과"},
// 		"patronType":{"id":2,"name":"대학원"},
// 		"memberNo":"2018121795"
// 	}
// }

type parentDept struct {
	// {
	// 	"id":65,
	// 	"code":"H0000476",
	// 	"name":"대학원"
	// },
}

type patronState struct {
	// :{
	// 	"id":8,
	// 	"name":"재학"
	// },
}

type branch struct {
	// :{
	// 	"id":9,
	// 	"name":"ERICA학술정보관",
	// 	"alias":"ERICA",
	// 	"libraryCode":"241050"
	// },
}

type dept struct {
	// :{
	// 	"id":653,
	// 	"code":"H0000706",
	// 	"name":"컴퓨터공학과"
	// },
}
type patronType struct {
	// {"id":2,"name":"대학원"},
}

type userLibChargeInfo struct {
	//
}

// Request
// curl -H 'Content-Type: application/json;charset=UTF-8'
// 	-XPOST 'http://lib.hanyang.ac.kr/pyxis-api/api/login'
// 	-d '{"loginId": "----",  "password": "----"}' -i
func (s *service) Login(id string, pw string) (*library.LoginInfo, error) {

	loginInfo := &library.LoginInfo{
		LoginId:  id,
		Password: pw,
	}
	// loginInfo, err := p.libraryService.Login(data.loginId, data.password)

	loginInfoRequest := loginRequest{id, pw}
	loginInfoBytes, _ := json.Marshal(loginInfoRequest)
	buff := bytes.NewBuffer(loginInfoBytes)

	loginReq, err := http.NewRequest("POST", defaultLibraryAddress+"/api/login", buff)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to make POST request"+defaultLibraryAddress+"/api/login"))
	}

	loginReq.Header.Add("Content-Type", "application/json;charset=UTF-8")

	log.Println(loginReq)

	loginClient := &http.Client{}
	loginResp, err := loginClient.Do(loginReq)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to complete POST request /pyxis-api/api/login"))
	}
	defer loginResp.Body.Close()

	log.Println(loginResp.Header)
	for _, cookie := range loginResp.Cookies() {
		if cookie.Name == "JSESSIONID" {
			loginInfo.JSessionID = cookie.Value
		}
	}
	log.Println("PRINT loginInfo : ", loginInfo)

	loginRespBody, err := ioutil.ReadAll(loginResp.Body)
	var response response
	if err := json.Unmarshal(loginRespBody, &response); err != nil {
		log.Println(errors.Wrap(err, "failed to unmarshal /pyxis-api/api/login"))
	}

	if response.success == false {
		log.Println("library login response return false")
		return nil, err
	}

	// Request
	// curl -H 'Content-Type: application/json;charset=UTF-8'
	// 	-XGET 'http://lib.hanyang.ac.kr/pyxis-api//1/api/charges?max=1000'
	// 	-c 'JSESSIONID='
	userLibReq, err := http.NewRequest("GET", defaultLibraryAddress+"/1/api/charges?max=1000", nil)
	if err != nil {
		return nil, err
	}

	userLibReq.AddCookie(&http.Cookie{Name: "JESSIONID", Value: loginInfo.JSessionID})

	http.Handle(defaultLibraryAddress+"/api/login", new(staticHandler))
	if err != nil {
		return nil, err
	}

	userLibClient := &http.Client{}
	userLibResp, err := loginClient.Do(userLibReq)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to complete GET request /pyxis-api/1/api/charges?max=1000"))
	}
	defer userLibResp.Body.Close()

	userLibRespBody, err := ioutil.ReadAll(userLibResp.Body)
	var userLibChargeInfo response
	if err := json.Unmarshal(userLibRespBody, &response); err != nil {
		log.Println(errors.Wrap(err, "failed to unmarshal /pyxis-api/1/api/charges?max=1000"))
	}

	return loginInfo, nil
}

func getURLHeaders(url string) map[string]interface{} {
	response, err := http.Head(url)
	if err != nil {
		log.Fatal("Error: Unable to download URL (", url, ") with error: ", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatal("Error: HTTP Status = ", response.Status)
	}

	headers := make(map[string]interface{})

	for k, v := range response.Header {
		headers[strings.ToLower(k)] = string(v[0])
	}

	return headers
}

func getURLHeaderByKey(url string, key string) string {

	headers := getURLHeaders(url)
	key = strings.ToLower(key)

	if value, ok := headers[key]; ok {
		return value.(string)
	}

	return ""
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

// IntialStore stores all login information at repository.
func (s *service) intialStore(fpath string) error {
	lines, err := readLines(fpath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read lines from loginInfo path"))
	}
	log.Println("initial loginInfo store started")
	for _, line := range lines {
		splitLine := strings.Split(line, "\t")
		id, pw := splitLine[0], splitLine[1]

		newLoginInfo := &library.LoginInfo{
			// TO BE IMPLEMENTED: kakao userkey
			UserKey:  "sample",
			LoginId:  id,
			Password: pw,
		}
		s.libraryLoginList.Store(newLoginInfo)
	}
	return nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
