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
	GetAllBirthday() ([]*birthday.Birthday, error)
	IntialStore(fpath string) error
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

// GetAllBirthday finds all birthday list in repository and returns it.
func (s *service) GetAllBirthday() ([]*birthday.Birthday, error) {
	resBirthday, err := s.birthdayList.FindAll()
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get birthday list"))
		return nil, err
	}
	return resBirthday, nil
}

// IntialStore stores birthday list in repository.
func (s *service) IntialStore(fpath string) error {
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
func NewService(r birthday.Repository) Service {
	return &service{
		birthdayList: r,
	}
}
