package generator

import (
	entity "Referral-System/generator/entity"
	"time"
)

//Repository interface
type Repository interface {
	Create(e *entity.Generator) (err error)
	FetchPassword(GeneratorID string) (Password string, err error)
	FetchReferralLink(GeneratorID string) (ReferralLink string, err error)
	UpdateExpirationTime(ReferralLink string, ExpirationTime time.Time) (err error)
}

//GRPCDriver interface
type GRPCDriver interface {
	Create(e *entity.Token) (AccessToken string, err error)
	Introspect(AccessToken string) (Role, ReferralLink string, err error)
}

//UseCase interface
type UseCase interface {
	CreateGenerator(ID, Name, Email, Password string) (GeneratedLink string, err error)
	LoginGenerator(ID, Password string) (AccessToken string, err error)
	ExtendTime(AccessToken string) (ExpirationTime string, err error)
}
