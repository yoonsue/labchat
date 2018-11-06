package location

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	locationModel "github.com/yoonsue/labchat/model/location"
)

// Service declares the methods that phone service provides.
type Service interface {
	GetLocation(string) ([]*locationModel.Location, error)
}

type service struct {
	locationList locationModel.Repository
}

// GetLocation finds loacation in repository and returns it.
func (s *service) GetLocation(request string) ([]*locationModel.Location, error) {
	resLocation, err := s.locationList.Find(request)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get phone number"))
		return nil, err
	}
	return resLocation, nil
}

// IntialStore stores all location at repository.
func (s *service) intialStore(fpath string) error {
	// TO BE IMPLEMENTED:
	// 1. store at the repository
	// 2. where to put this fuction(maybe NewService)
	lines, err := readLines(fpath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read lines from phone path"))
	}
	log.Println("initial location store started")
	for _, line := range lines {
		name := ""
		location := ""
		if strings.HasPrefix(line, "=") {
			name += line
		} else {
			splitLine := strings.Split(line, "\t")
			name, location = splitLine[0], splitLine[1]
			// extenInt, err := strconv.Atoi(exten)
			if err != nil {
				log.Println("exten is not int type")
			}
			newLocation := &locationModel.Location{
				Name:     name,
				Location: location,
			}
			s.locationList.Store(newLocation)
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
func NewService(r locationModel.Repository, fpath string) Service {
	s := &service{
		locationList: r,
	}
	s.intialStore(fpath)
	return s
}
