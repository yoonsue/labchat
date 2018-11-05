package phone

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/yoonsue/labchat/model/phone"
	"github.com/yoonsue/labchat/repository/inmem"
)

func TestGetPhone(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "tmpPhone")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write([]byte("department1\t0001\ndepartment2\t1010\ndepartment3\t2020\n"))
	if err != nil {
		t.Fatal(err)
	}

	r := inmem.NewPhoneRepository()
	s := NewService(r, tmpFile.Name())

	testCases := []struct {
		depSubString phone.Department
		expected     []*phone.Phone
	}{
		{
			"de",
			[]*phone.Phone{
				{Department: "department1", Extension: "0001"},
				{Department: "department2", Extension: "1010"},
				{Department: "department3", Extension: "2020"},
			},
		},
		{
			"department1",
			[]*phone.Phone{
				{Department: "department1", Extension: "0001"},
			},
		},
	}
	for _, c := range testCases {
		gotPhoneList, err := s.GetPhone(string(c.depSubString))
		if err != nil {
			t.Error("GetPhone failed")
		}

		if !reflect.DeepEqual(c.expected, gotPhoneList) {
			t.Error("NewService error")
		}
	}
}

func TestReadLines(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "tmpPhone")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write([]byte("line1\nline2\tcontent\n"))
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"line1", "line2\tcontent"}
	gotLines, err := readLines(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, gotLines) {
		t.Errorf("expected Lines differs with gotten one")
	}
}

func TestNewService(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "tmpBirth")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write([]byte("department1\t0001\ndepartment2\t1010\ndepartment3\t2020\n"))
	if err != nil {
		t.Fatal(err)
	}

	r := inmem.NewPhoneRepository()
	gotService := NewService(r, tmpFile.Name())

	s := &service{
		phonebook: r,
	}

	expected, _ := s.GetPhone("department1")
	got, _ := gotService.GetPhone("department1")
	if !reflect.DeepEqual(expected, got) {
		t.Error("NewService error")
	}
}
