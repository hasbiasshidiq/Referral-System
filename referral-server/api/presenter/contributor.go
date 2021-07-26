package presenter

// Contribution presenter data
type Contribution struct {
	Status string `json:"status"`
}

type Contributor struct {
	Email        string `json:"email"`
	ReferralLink string `json:"referral_link"`
	Contribution string `json:"contribution"`
}
