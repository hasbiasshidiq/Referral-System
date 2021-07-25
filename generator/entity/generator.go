package entity

import (
	"math/rand"
	"strings"
	"time"

	"Referral-System/generator/config"
)

// Generator data
type Generator struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	GeneratedLink string `json:"generated_link"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ExpirateAt    time.Time
}

// NewGenerator is a function to create new generator data
func NewGenerator(ID, Name, Email, Password string) (*Generator, error) {
	gen := &Generator{
		ID:       ID,
		Name:     Name,
		Email:    Email,
		Password: Password,

		GeneratedLink: GenerateRandomLink(40),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ExpirateAt:    time.Now().AddDate(0, 0, 7),
	}
	err := gen.Validate()
	if err != nil {
		return nil, err
	}
	return gen, nil
}

//Validate validate Generator element
func (g *Generator) Validate() error {
	if g.ID == "" || g.Name == "" || g.Email == "" || g.Password == "" || g.GeneratedLink == "" {
		return ErrInvalidEntity
	}
	return nil
}

// GenerateRandomLink will generate random link with a given last subset length
func GenerateRandomLink(length int) (result string) {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	result = config.SHARED_LINK_ENDPOINT + b.String()

	return
}
