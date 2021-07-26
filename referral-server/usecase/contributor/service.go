package contributor

import (
	"log"
	entity "referral-server/entity"
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

//ListContributor get all contributor for specific generator / referrall link
func (s *Service) ListContributor(AccessToken string) ([]*entity.Contributor, error) {

	// first check if token is valid
	role, referralLink, err := s.grpc.Introspect(AccessToken)
	if err != nil {
		log.Println("error ListContributor usecase - IntrospectToken err : ", err)
		return nil, err
	}

	if role != "generator" {
		err = entity.ErrUnauthorizedAccess
		return nil, err
	}

	contributors, err := s.repo.List(referralLink)
	if err != nil {
		return nil, err
	}
	if len(contributors) == 0 {
		return nil, entity.ErrNotFound
	}
	return contributors, nil
}
