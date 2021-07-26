package entity

import (
	"time"
)

// Token data
type Token struct {
	Issuer string `json:"iss"`

	Iat time.Time `json:"iat"`
	Exp time.Time `json:"exp"`

	// Custom token claim for authorization
	ReferralLink string `json:"referral_link"`
	Role         string `json:"role"`
}

// NewToken is a function to create new Token data
func NewToken(ReferralLink, Role string, Exp time.Time) (*Token, error) {

	tok := &Token{
		Issuer:       "Referral-System",
		Iat:          time.Now(),
		Exp:          Exp,
		ReferralLink: ReferralLink,
		Role:         Role,
	}
	err := tok.Validate()
	if err != nil {
		return nil, err
	}
	return tok, nil
}

//Validate validate Token element
func (t *Token) Validate() error {
	if time.Now().After(t.Exp) {
		return ErrInvalidExpirationTime
	}

	if t.Issuer == "" || t.ReferralLink == "" || t.Role == "" {
		return ErrInvalidEntity
	}
	return nil
}
