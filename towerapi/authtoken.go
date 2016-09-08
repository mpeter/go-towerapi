package towerapi

import "github.com/dghubble/sling"

const authTokenBasePath = "authtoken/"

// AuthTokenService is an interface for interfacing with the AuthToken
// endpoints of the Ansible Tower API
type AuthTokenService struct {
	sling *sling.Sling
}

// NewAuthTokenService handles communication with auth token related methods of the
// Ansible Tower API.
func NewAuthTokenService(sling *sling.Sling) *OrganizationsService  {
	return &AuthTokenService{
		sling: sling.Path(authTokenBasePath),
	}
}

// AuthToken represents a Ansible Tower AuthToken.
type AuthToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

// AuthTokenCreateRequest represents a request to create a new key.
type AuthTokenCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
