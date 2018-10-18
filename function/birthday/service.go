package birthday

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model/birthday"
)

// Service declares the methods that phone service provides.
type Service interface {
	GetBirthday(name string) (*birthday.Birthday, error)
	CheckBirthday() []*birthday.Birthday
}

type service struct {
	birthdayList birthday.Repository
}

// GetBirthday finds birthday in repository and returns it.
func (s *service) GetBirthday(name string) (*birthday.Birthday, error) {
	resBirthday, err := s.birthdayList.Find(name)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get birthday from name"))
		return nil, err
	}
	return resBirthday, nil
}

// getAllBirthday finds all birthday list in repository and returns it.
func (s *service) getAllBirthday() ([]*birthday.Birthday, error) {
	resBirthday, err := s.birthdayList.FindAll()
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get birthday list"))
		return nil, err
	}
	return resBirthday, nil
}

// CheckBirthday finds the name list matching birhtday with today in repository and returns it.
func (s *service) CheckBirthday() []*birthday.Birthday {
	currentTime := time.Now().Local()
	currentTime = currentTime.Add(time.Hour * (-1) * 9)
	currentString := currentTime.Format("2006-01-02")

	var todayIsBirthList []*birthday.Birthday
	mapBirthday, _ := s.getAllBirthday()
	for _, tmp := range mapBirthday {
		tmpBirth := tmp.GetBirth()
		if (tmpBirth[:2] == currentString[5:7]) && (tmpBirth[6:8] == currentString[8:10]) {
			todayIsBirthList = append(todayIsBirthList, tmp)
		}
	}
	return todayIsBirthList
}

// IntialStore stores birthday list in repository.
func (s *service) intialStore(fpath string) error {
	lines, err := readLines(fpath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read lines from phone path"))
	}
	log.Println("initial birthday store started")
	for _, line := range lines {
		splitLine := strings.Split(line, "\t")
		name, birth := splitLine[0], splitLine[1]
		birthInt, err := strconv.Atoi(birth)
		if err != nil {
			log.Println("exten is not int type")
		}
		newBirthday := &birthday.Birthday{
			Name:        name,
			DateOfBirth: birthInt,
		}
		s.birthdayList.Store(newBirthday)
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

// NewService return struct which provides Service interface
func NewService(r birthday.Repository, fpath string) Service {
	s := &service{
		birthdayList: r,
	}
	s.intialStore(fpath)
	return s
}
