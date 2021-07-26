package entity_test

import (
	entity "referral-server/entity"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {

	exp := time.Now().AddDate(0, 0, 1)

	u, err := entity.NewToken("http://127.0.0.1:/portal/xyz", "generator", exp)

	assert.Nil(t, err)

	assert.Equal(t, u.Issuer, "Referral-System")
	assert.Equal(t, u.Exp.After(u.Iat), true)
	assert.Equal(t, u.ReferralLink, "http://127.0.0.1:/portal/xyz")
	assert.Equal(t, u.Role, "generator")

}

func TestValidateToken(t *testing.T) {
	type test struct {
		referralLink string
		role         string
		exp          time.Time
		want         error
	}

	tests := []test{
		{
			referralLink: "http://127.0.0.1:/portal/xyz",
			role:         "generator",
			exp:          time.Now().AddDate(0, 0, 1),

			want: nil,
		},
		{
			referralLink: "",
			role:         "generator",
			exp:          time.Now().AddDate(0, 0, 1),

			want: entity.ErrInvalidEntity,
		},
		{
			referralLink: "http://127.0.0.1:/portal/xyz",
			role:         "",
			exp:          time.Now().AddDate(0, 0, 1),

			want: entity.ErrInvalidEntity,
		},
		{
			referralLink: "http://127.0.0.1:/portal/xyz",
			role:         "generator",
			exp:          time.Now().AddDate(0, 0, -1),

			want: entity.ErrInvalidExpirationTime,
		},
	}

	for _, tc := range tests {

		_, err := entity.NewToken(tc.referralLink, tc.role, tc.exp)
		assert.Equal(t, err, tc.want)
	}

}
