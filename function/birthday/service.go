package birthday

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model/birthday"
)

// Service declares the methods that phone service provides.
type Service interface {
	GetBirthday(name string) (*birthday.Birthday, error)
	IntialStore(fpath string) error
}

type service struct {
	birthdayList birthday.Repository
}

// GetBirthday finds birthday in repository and returns it.
func (s *service) GetBirthday(name string) (*birthday.Birthday, error) {
	resBirthday, err := s.birthdayList.Find(name)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get phone number"))
		return nil, err
	}
	return resBirthday, nil
}

// IntialStore stores birthday list in repository.
func (s *service) IntialStore(fpath string) error {
	// TO BE IMPLEMENTED:
	// 1. store at the repository
	// 2. where to put this fuction(maybe NewService)
	lines, err := readLines(fpath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read lines from phone path"))
	}
	log.Println("initial birthday store started")
	for _, line := range lines {
		splitLine := strings.Split(line, "\t")
		name, info := splitLine[0], splitLine[1]
		age, birth := info[:2], info[2:]
		birth = birth[0:2] + "월 " + birth[2:] + "일"
		ageInt, err := strconv.Atoi(age)
		// birthInt, err := strconv.Atoi(birth)
		if err != nil {
			log.Println("exten is not int type")
		}
		newBirthday := &birthday.Birthday{
			Name:     name,
			Birthday: birth,
			Age:      (118 - ageInt + 1),
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
func NewService(r birthday.Repository) Service {
	return &service{
		birthdayList: r,
	}
}
