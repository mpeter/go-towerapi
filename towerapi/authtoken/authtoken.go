package authtoken

import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/base"
	"net/http"
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

func (s *Service) Create(r *CreateRequest) (*AuthToken, *http.Response, error) {
	var token *AuthToken
	var towerError *base.TowerError
	response, err := s.sling.New().Post("").BodyJSON(r).Receive(&token, &towerError)
	if err == nil {
		if towerError != nil {
			err = towerError
		}
	}
	return token, response, err
}
