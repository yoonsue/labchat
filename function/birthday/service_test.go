package birthday

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/yoonsue/labchat/model/birthday"
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
	cS := "2018-01-16"

	tmpFile, err := ioutil.TempFile("", "tmpBirth")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write([]byte("name1\t900116\nname2\t881116\nsue\t990116\n"))
	if err != nil {
		t.Fatal(err)
	}

	r := inmem.NewBirthdayRepository()
	s := NewService(r, tmpFile.Name())
	expected := []*birthday.Birthday{
		{Name: "name1", DateOfBirth: 900116},
		{Name: "name3", DateOfBirth: 990116},
	}

	gotBirth := s.CheckBirthday(cS)
	if reflect.DeepEqual(gotBirth, expected) {
		t.Errorf("expected []*birthday.Birthday differs with gotten one")
	}

}
func TestIntialStore(t *testing.T) {

}
