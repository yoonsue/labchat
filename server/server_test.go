package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/yoonsue/labchat/function/menu"
)

func TestNewServer(t *testing.T) {
	var ms menu.Service
	s := &Server{
		cfg: &Config{
			Address: "localhost:8080",
		},
		menuService: ms,
	}
	gotServer, _ := NewServer(s.cfg, ms)
	if s.cfg != gotServer.cfg {
		t.Errorf("expected %s, got %s", s.cfg, gotServer.cfg)
	}
	if s.menuService != gotServer.menuService {
		t.Errorf("expected %s, got %s", s.menuService, gotServer.menuService)
	}
}

// func TestStart(t *testing.T) {
// 	var s *Server
// 	rou := s.Router()

// }
func TestRouter(t *testing.T) {

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
	expected := "{\"type\":\"text\",\"buttons\":null}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMessageHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/message", strings.NewReader(""))
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
	expected := "{\"message\":{\"text\":\"....????\"}}"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestFriendHandler(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/friend", strings.NewReader("{\"user_key\" : \"HASHED_USER_KEY\" }"))
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
	expected := "Hello~"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// type Service interface {
// 	GetSchool(url string) *model.Menu
// }
// type fakeService struct {
// }
// func (s fakeService) GetSchool(url string) *model.Menu {
// 	return &model.Menu{Restaurant: "restA", TodayMenu: "menuA"}
// }

// func TestMsgFor(t *testing.T) {
// 	var s *Server
// 	// s := &Server{
// 	// 	cfg: &Config{
// 	// 		Address: "localhost:8080",
// 	// 	},
// 	// }

// 	testCases := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{
// 			"status",
// 			"TIME : ",
// 		},
// 		{
// 			"menu",
// 			"restA",
// 		},
// 		{
// 			"hello",
// 			"hello....????",
// 		},
// 	}

// 	for _, c := range testCases {
// 		result := s.msgFor(c.input)
// 		if !strings.HasPrefix(result, c.expected) {
// 			t.Errorf("start with '%s' on input %s, got %s", c.expected, c.input, result)
// 		}
// 	}
// }
