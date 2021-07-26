package referral

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

// CreateContributorToken create a Token
func (s *Service) CreateContributorToken(ReferralLink string) (AccessToken string, err error) {

	exp, err := s.repo.FetchExpirationTime(ReferralLink)
	if err != nil {
		log.Println("CreateContributorToken err : ", err.Error())
		return
	}
	t, err := entity.NewToken(ReferralLink, "contributor", exp)
	if err != nil {
		log.Println("CreateContributorToken err : ", err.Error())
		return
	}

	accessTok, err := s.grpc.Create(t)

	return accessTok, err
}
