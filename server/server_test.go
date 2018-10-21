package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/yoonsue/labchat/function/birthday"
	"github.com/yoonsue/labchat/function/menu"
	"github.com/yoonsue/labchat/function/phone"
	"github.com/yoonsue/labchat/function/status"
	"github.com/yoonsue/labchat/repository/inmem"
)

func TestNewServer(t *testing.T) {
	var ms menu.Service
	var ps phone.Service
	var ss status.Service
	var bs birthday.Service
	s := Server{
		cfg: &Config{
			Address: "localhost:8080",
		},
		menuService:     ms,
		phoneService:    ps,
		statusService:   ss,
		birthdayService: bs,
	}
	gotServer, _ := NewServer(s.currentTime, s.cfg, ms, ps, ss, bs)
	if s.cfg != gotServer.cfg {
		t.Errorf("expected %s, got %s", s.cfg, gotServer.cfg)
	}
	if s.menuService != gotServer.menuService {
		t.Errorf("expected %s, got %s", s.menuService, gotServer.menuService)
	}
}

func TestStart(t *testing.T) {
	var ms menu.Service
	var ps phone.Service
	var ss status.Service
	var bs birthday.Service

	s := Server{
		cfg: &Config{
			Address: "localhost:8080",
		},
		menuService:     ms,
		phoneService:    ps,
		statusService:   ss,
		birthdayService: bs,
	}
	gotMux := s.Start()

	testCases := []struct {
		method   string
		uri      string
		query    io.Reader
		expected string
	}{
		{
			"GET", "/keyboard",
			nil,
			"{\"type\":\"buttons\",\"buttons\":[\"도움말\",\"시작하기\"]}\n",
		},
		{
			"POST", "/message",
			strings.NewReader("{\"user_key\" : \"HASHED_USER_KEY\", \"type\":\"text\",\"content\":\"?\" }"),
			"{\"message\":{\"text\":\"?....????\"}}",
		},
		{
			"POST", "/friend",
			strings.NewReader("{\"user_key\" : \"HASHED_USER_KEY\" }"),
			"Hello, my friend",
		},
	}
	for _, c := range testCases {
		req, _ := http.NewRequest(c.method, c.uri, c.query)
		res := httptest.NewRecorder()
		gotMux.ServeHTTP(res, req)

		if res.Body.String() != c.expected {
			t.Errorf("Expected hello %s but got %s", c.expected, res.Body.String())
		}
	}
}
func TestMessageKey(t *testing.T) {
	testCases := []struct {
		key      string
		expected string
	}{
		{
			"hi",
			"hello", // TestCase 1
		},
		{
			"hello",
			"hi",
		},
		{
			"name",
			"LABchat",
		},
		{
			"no",
			"none",
		},
	}

	for _, c := range testCases {
		result := messageKey(c.key)
		if result != c.expected {
			t.Errorf("expected %s for key %s, got %s", c.expected, c.key, result)
		}
	}
}

var s *Server

func TestKeyboardHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/keyboard", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.keyboardHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{\"type\":\"buttons\",\"buttons\":[\"도움말\",\"시작하기\"]}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMessageHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/message", strings.NewReader("{\"user_key\" : \"HASHED_USER_KEY\", \"type\":\"text\",\"content\":\"?\" }"))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.messageHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{\"message\":{\"text\":\"?....????\"}}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestFriendHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/friend", strings.NewReader("{\"user_key\" : \"USER_KEY\" }"))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.friendHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Hello, my friend"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestFriendDeleteHandler(t *testing.T) {
	baseReq, err := http.NewRequest("POST", "/friend", strings.NewReader("{\"user_key\" : \"USER_KEY\" }"))

	// Create a request to pass to our handler.
	req, err := http.NewRequest("DELETE", "/friend/USER_KEY", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	baseHandler := http.HandlerFunc(s.friendHandler)
	baseHandler.ServeHTTP(rr, baseReq)
	handler := http.HandlerFunc(s.friendDeleteHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Hello, my friend\nuser deleted"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestChatroomDeleteHandler(t *testing.T) {
	baseReq, err := http.NewRequest("POST", "/friend", strings.NewReader("{\"user_key\" : \"HASHED_USER_KEY\" }"))

	// Create a request to pass to our handler.
	req, err := http.NewRequest("DELETE", "/chat_room/USER_KEY", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	baseHandler := http.HandlerFunc(s.friendHandler)
	baseHandler.ServeHTTP(rr, baseReq)
	handler := http.HandlerFunc(s.chatroomDeleteHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Hello, my friend\nchatroom deleted"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMsgFor(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "tmpBirth")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Temp file created:", tmpFile.Name())

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err = tmpFile.Write([]byte("조윤수\t960116\n")); err != nil {
		t.Fatal(err)
	}

	mr := inmem.NewMenuRepository()
	ms := menu.NewService(mr)
	var ss status.Service
	br := inmem.NewBirthdayRepository()
	bs := birthday.NewService(br, tmpFile.Name())

	s := Server{
		currentTime: "0000-00-00",
		cfg: &Config{
			Address: "localhost:8080",
		},
		menuService:     ms,
		statusService:   ss,
		birthdayService: bs,
	}

	testCases := []struct {
		input    []string
		expected string
	}{
		{
			strings.Fields("도움말"),
			"LABchat",
		},
		{
			strings.Fields("시작하기"),
			"필요한",
		},
		// {
		// 	strings.Fields("status"),
		// 	"TEMP",
		// }, test
		{
			strings.Fields("menu"),
			"교직원식당",
		},
		{
			strings.Fields("phone"),
			"no department",
		},
		{
			strings.Fields("생일"),
			"no name",
		},
		{
			strings.Fields("hello"),
			"hello....????",
		},
	}

	for _, c := range testCases {
		result := s.msgFor(c.input)
		if !strings.HasPrefix(result, c.expected) {
			t.Errorf("start with '%s' on input %s, got %s", c.expected, c.input, result)
		}
	}
}
