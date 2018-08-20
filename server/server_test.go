package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func (s *Server) TestKeyboardHandler(t *testing.T) {
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
	expected := `{"type":"text","buttons":null}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func TestMsgFor(t *testing.T) {
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
// 			"\n==교직원식당==\n",
// 		},
// 		{
// 			"hello",
// 			"hello....????",
// 		},
// 	}

// 	for _, c := range testCases {
// 		result := msgFor(c.input)
// 		if !strings.HasPrefix(result, c.expected) {
// 			t.Errorf("start with '%s' on input %s, got %s", c.expected, c.input, result)
// 		}
// 	}
// }
