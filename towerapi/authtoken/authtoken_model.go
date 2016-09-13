package authtoken

// AuthToken represents a Ansible Tower AuthToken.
type AuthToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

// AuthTokenCreateRequest represents a request to create a new key.
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
