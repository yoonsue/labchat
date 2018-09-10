package phone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model/phone"
)

// Service declares the methods that phone service provides.
type Service interface {
	GetPhone(department phone.Department) (*phone.Phone, error)
	IntialStore(fpath string) error
}

type service struct {
	phonebook phone.Repository
}

// GetPhone finds phone number in repository and returns it.
func (s *service) GetPhone(department phone.Department) (*phone.Phone, error) {
	resPhone, err := s.phonebook.Find(department)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get phone number"))
		return nil, err
	}
	return resPhone, nil
}

// IntialStore stores all phone in repository.
func (s *service) IntialStore(fpath string) error {
	// TO BE IMPLEMENTED:
	// 1. store at the repository
	// 2. where to put this fuction(maybe NewService)
	lines, err := readLines(fpath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read lines from phone path"))
	}
	log.Println("initial store started")
	for _, line := range lines {
		dept := ""
		exten := ""
		fmt.Println(line)
		if strings.HasPrefix(line, "=") {
			dept += line
		} else {
			splitLine := strings.Split(line, "\t")
			dept, exten = splitLine[0], splitLine[1]
			extenInt, err := strconv.Atoi(exten)
			if err != nil {
				log.Println("exten is not int type")
			}
			newPhone := &phone.Phone{
				Department: phone.Department(dept),
				Extension:  extenInt,
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
func NewService(r phone.Repository) Service {
	return &service{
		phonebook: r,
	}
}
