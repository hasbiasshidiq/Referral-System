package contributor

import (
	entity "Referral-System/generator/entity"
	"log"
)

// Service for Token usecase
type Service struct {
	repo Repository
	grpc GRPCDriver
}

// NewService create new service
func NewService(r Repository, g GRPCDriver) *Service {
	return &Service{
		repo: r,
		grpc: g,
	}
}

// Contribute create a contribution
func (s *Service) Contribute(Email, AccessToken string) (err error) {

	role, referralLink, err := s.grpc.Introspect(AccessToken)
	if err != nil {
		return
	}

	if role != "contributor" {
		err = entity.ErrUnauthorizedAccess
		return
	}

	c, err := entity.NewContributor(Email, referralLink)
	if err != nil {
		log.Println("CreateContributor err : ", err.Error())
		return
	}

	err = s.repo.Contribute(c)

	return
}
