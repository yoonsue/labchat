package birthday

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// Birthday = VO
type Birthday struct {
	Name        string
	DateOfBirth int
}

// Repository declares the methods that repository provides.
type Repository interface {
	Find(key string) (*Birthday, error)
	FindAll() ([]*Birthday, error)
	Store(key *Birthday) error
	Clean() error
}

// GetAge returns the Age from given year of birth.
func (b Birthday) GetAge() int {
	currentTime := time.Now().Local()
	currentString := currentTime.Format("2006-01-02")
	currentYear, err := strconv.Atoi(currentString[1:4])
	if err != nil {
		log.Println("year is not int type")
	}

	year := b.DateOfBirth / 10000
	return (100 + currentYear - year + 1)
}

// GetBirth returns the birth as form "00월 00일".
func (b Birthday) GetBirth() string {
	month := (b.DateOfBirth % 10000) / 100
	strMonth := fmt.Sprintf("%02d", month)
	day := b.DateOfBirth % 100
	strDay := fmt.Sprintf("%02d", day)
	return strMonth + "월 " + strDay + "일"
}
