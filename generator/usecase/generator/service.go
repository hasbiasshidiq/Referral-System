package generator

import (
	entity "Referral-System/generator/entity"
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Service for Generator usecase
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

// LoginGenerator create a Generator
func (s *Service) LoginGenerator(ID, Password string) (AccessToken string, err error) {

	p, err := s.repo.FetchPassword(ID)
	if err == sql.ErrNoRows {
		err = entity.ErrInvalidCredentials
		return
	}

	if err != nil {
		log.Println("LoginGenerator UseCase - FetchPassword err : ", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(p), []byte(Password))
	if err != nil {
		log.Println("LoginGenerator UseCase - CompareHash err : ", err)
		err = entity.ErrInvalidCredentials
		return
	}

	referralLink, err := s.repo.FetchReferralLink(ID)
	if err != nil {
		log.Println("LoginGenerator UseCase - Fetch Referral Link err : ", err)
		return
	}

	// token will be expired in one day
	tokenExpirationTime := time.Now().AddDate(0, 0, 1)

	t, err := entity.NewToken(referralLink, "generator", tokenExpirationTime)
	if err != nil {
		log.Println("LoginGenerator UseCase - NewToken err : ", err)
		return
	}

	accessTok, err := s.grpc.Create(t)

	return accessTok, err

}

// ExtendTime extend time of specific link
func (s *Service) ExtendTime(AccessToken string) (ExpirationTime string, err error) {

	// first check if token is valid
	role, referralLink, err := s.grpc.Introspect(AccessToken)
	if err != nil {
		log.Println("error ExtendTime usecase - IntrospectToken err : ", err)
		return
	}

	if role != "generator" {
		err = entity.ErrUnauthorizedAccess
		return
	}

	Exp := time.Now().AddDate(0, 0, 7)
	err = s.repo.UpdateExpirationTime(referralLink, Exp)
	if err != nil {
		log.Println("error ExtendTime usecase - UpdateExpirationTime err : ", err)
		return
	}

	ExpirationTime = Exp.Format("2006-01-02 15:04:05")

	return

}
