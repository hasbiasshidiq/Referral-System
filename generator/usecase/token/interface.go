package token

import (
	entity "Referral-System/generator/entity"
	"time"
)

//Repository interface
type Repository interface {
	FetchExpirationTime(ReferralLink string) (Exp time.Time, err error)
}

//GRPCDriver interface
type GRPCDriver interface {
	Create(e *entity.Token) (AccessToken string, err error)
}

//UseCase interface
type UseCase interface {
	CreateContributorToken(ReferralLink string) (AccessToken string, err error)
}
