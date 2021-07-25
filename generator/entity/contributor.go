package entity

// Contributor data
type Contributor struct {
	Email        string `json:"email"`
	ReferralLink string `json:"referral_link"`
}

// NewContributor is a function to create new contributor data
func NewContributor(Email, ReferralLink string) (*Contributor, error) {
	cont := &Contributor{

		Email:        Email,
		ReferralLink: ReferralLink,
	}
	err := cont.Validate()
	if err != nil {
		return nil, err
	}
	return cont, nil
}

//Validate validate Contributor element
func (c *Contributor) Validate() error {
	if c.Email == "" || c.ReferralLink == "" {
		return ErrInvalidEntity
	}
	return nil
}
