package authtoken

import (
	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/sling"
)

const basePath = "authtoken/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth token related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.New().Path(basePath),
	}
}

// Create passes credentials and returns a token
func (s *Service) Create(r *CreateRequest) (*AuthToken, error) {
	token := new(AuthToken)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(token, apierr)
	return token, errors.BuildError(err, apierr)
}
