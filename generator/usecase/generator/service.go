package generator

import (
	entity "Referral-System/generator/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Service for Generator usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateGenerator create a Generator
func (s *Service) CreateGenerator(ID, Name, Email, Password string) (GeneratedLink string, err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	a, err := entity.NewGenerator(ID, Name, Email, string(hashedPassword))
	if err != nil {
		return "", err
	}

	err = s.repo.Create(a)

	return a.GeneratedLink, err
}
