package entity_test

import (
	entity "referral-server/entity"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenerator(t *testing.T) {

	u, err := entity.NewGenerator("ronaldinho", "Ronaldinho", "rdbrazil@gmail.com", "not_a_secret_pass")
	assert.Nil(t, err)

	assert.Equal(t, u.ID, "ronaldinho")
	assert.Equal(t, u.Name, "Ronaldinho")
	assert.Equal(t, u.Email, "rdbrazil@gmail.com")
	assert.Equal(t, u.Password, "not_a_secret_pass")

	assert.NotNil(t, u.GeneratedLink)
	assert.NotNil(t, u.CreatedAt)
	assert.NotNil(t, u.UpdatedAt)
	assert.NotNil(t, u.ExpirateAt)

}

func TestValidateGenerator(t *testing.T) {
	type test struct {
		id       string
		name     string
		email    string
		password string
		want     error
	}

	tests := []test{
		{
			id:       "jesse",
			name:     "Jesse Lingard",
			email:    "jelingz@gmail.com",
			password: "pass",
			want:     nil,
		},
		{
			id:       "",
			name:     "Jesse Lingard",
			email:    "jelingz@gmail.com",
			password: "pass",
			want:     entity.ErrInvalidEntity,
		},
		{
			id:       "jesse",
			name:     "",
			email:    "jelingz@gmail.com",
			password: "pass",
			want:     entity.ErrInvalidEntity,
		},
		{
			id:       "jesse",
			name:     "Jesse Lingard",
			email:    "",
			password: "pass",
			want:     entity.ErrInvalidEntity,
		},
		{
			id:       "jesse",
			name:     "Jesse Lingard",
			email:    "jelingz@gmail.com",
			password: "",
			want:     entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {

		_, err := entity.NewGenerator(tc.id, tc.name, tc.email, tc.password)
		assert.Equal(t, err, tc.want)
	}

}
