package phone

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model/phone"
)

// Service declares the methods that phone service provides.
type Service interface {
	GetPhone(department phone.Department) ([]*phone.Phone, error)
}

type service struct {
	phonebook phone.Repository
}

func (s *service) GetPhone(request string) ([]*phone.Phone, error) {
	if _, err := strconv.Atoi(request); err == nil {
		number, _ := strconv.Atoi(request)
		p, _ := s.getPhoneByNumber(number)
	} else {
	}
}

// GetPhone finds phone number in repository and returns it.
func (s *service) getPhoneByDepartment(department phone.Department) ([]*phone.Phone, error) {
	resPhone, err := s.phonebook.Find(department)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get phone number"))
		return nil, err
	}
	return resPhone, nil
}

// GetPhone finds phone number in repository and returns it.
func (s *service) getPhoneByNumber(number int) ([]*phone.Phone, error) {
	resPhone, err := s.phonebook.Find(number)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get phone number"))
		return nil, err
	}
	return resPhone, nil
}

// IntialStore stores all phone in repository.
func (s *service) intialStore(fpath string) error {
	// TO BE IMPLEMENTED:
	// 1. store at the repository
	// 2. where to put this fuction(maybe NewService)
	lines, err := readLines(fpath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read lines from phone path"))
	}
	log.Println("initial phone store started")
	for _, line := range lines {
		dept := ""
		exten := ""
		if strings.HasPrefix(line, "=") {
			dept += line
		} else {
			splitLine := strings.Split(line, "\t")
			dept, exten = splitLine[0], splitLine[1]
			// extenInt, err := strconv.Atoi(exten)
			if err != nil {
				log.Println("exten is not int type")
			}
			newPhone := &phone.Phone{
				Department: phone.Department(dept),
				Extension:  exten,
			}
			s.phonebook.Store(newPhone)
		}
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
func NewService(r phone.Repository, fpath string) Service {
	s := &service{
		phonebook: r,
	}
	s.intialStore(fpath)
	return s
}
