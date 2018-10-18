package birthday

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/yoonsue/labchat/repository/inmem"
)

func TestGetBirthday(t *testing.T) {

	tmpFile, err := ioutil.TempFile("", "tmpBirth")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write([]byte("조윤수\t960116\n"))
	if err != nil {
		t.Fatal(err)
	}

	r := inmem.NewBirthdayRepository()
	s := NewService(r, tmpFile.Name())
	testCases := []struct {
		name         string
		expectedDate int
	}{
		{
			"조윤수",
			960116,
		},
		// {
		// 	"no",
		// 	10,
		// },
	}
	for _, c := range testCases {
		gotBirth, _ := s.GetBirthday(c.name)
		// if gotBirth == nil {
		// 	t.Errorf("nil\n")
		// } else
		if c.expectedDate != gotBirth.DateOfBirth {
			t.Errorf("expected %d, got %d", c.expectedDate, gotBirth.DateOfBirth)
		}
	}

}
func TestCheckBirthday(t *testing.T) {

}
func TestIntialStore(t *testing.T) {

}
