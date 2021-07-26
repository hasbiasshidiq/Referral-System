package contributor

import entity "referral-server/entity"

//Repository interface
type Repository interface {
	Contribute(e *entity.Contributor) (err error)
	List(referralLink string) (contributors []*entity.Contributor, err error)
}

//GRPCDriver interface
type GRPCDriver interface {
	Introspect(AccessToken string) (Role, ReferralLink string, err error)
}

//UseCase interface
type UseCase interface {
	Contribute(Email, AccessToken string) (err error)
	ListContributor(AccessToken string) ([]*entity.Contributor, error)
}
