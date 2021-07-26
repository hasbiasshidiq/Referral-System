package entity_test

import (
	entity "referral-server/entity"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContributor(t *testing.T) {

	u, err := entity.NewContributor("ronaldo@yahoo.co.id", "http://localhost:8080/access/xyu")
	assert.Nil(t, err)

	assert.Equal(t, u.Email, "ronaldo@yahoo.co.id")
	assert.Equal(t, u.ReferralLink, "http://localhost:8080/access/xyu")

}

func TestValidateContributor(t *testing.T) {
	type test struct {
		email        string
		referralLink string
		want         error
	}

	tests := []test{
		{
			email:        "jelingz@gmail.com",
			referralLink: "http://localhost:8080/access/xyu",
			want:         nil,
		},
		{
			email:        "",
			referralLink: "http://localhost:8080/access/xyu",
			want:         entity.ErrInvalidEntity,
		},
		{
			email:        "jelingz@gmail.com",
			referralLink: "",
			want:         entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {

		_, err := entity.NewContributor(tc.email, tc.referralLink)
		assert.Equal(t, err, tc.want)
	}

}
