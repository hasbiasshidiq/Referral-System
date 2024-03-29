package presenter

// Generator presenter data
type Generator struct {
	GeneratedLink string `json:"generated_link"`
}

// LoginGenerator presenter data
type LoginGenerator struct {
	AccessToken string `json:"access_token"`
}

// ExtendingTime presenter data
type ExtendingTime struct {
	Status string `json:"status"`
}
