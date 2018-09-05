package phone

import (
	"log"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model/phone"
)

// Service declares the methods that phone service provides.
type Service interface {
	GetPhone(department phone.Department) (*phone.Phone, error)
}

type service struct {
	phonebook phone.Repository
}

func (s *service) GetPhone(department phone.Department) (*phone.Phone, error) {
	resPhone, err := s.phonebook.Find(department)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get phone number"))
		return nil, err
	}
	return resPhone, nil
}

func (s *service) intialStore() {
	// TO BE IMPLEMENTED:
	// 1. store at the repository
	// 2. where to put this fuction(maybe NewService)
}

// NewService return struct which provides Service interface
func NewService(r phone.Repository) Service {
	return &service{
		phonebook: r,
	}
}
