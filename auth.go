package godi_example

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

// AuthToken holds authentication token details with refresh token
type AuthToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
