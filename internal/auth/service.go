package auth

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthenticatedUser(username, password string) (*User, error) {

	return nil, nil
}
