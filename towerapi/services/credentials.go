package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const credentialsBasePath = "credentials/"

type CredentialsService struct {
	sling *sling.Sling
}

func NewCredentialsService(sling *sling.Sling) *CredentialsService {
	return &CredentialsService {
		sling: sling.Path(credentialsBasePath),
	}
}

// List all credentials
func (s *CredentialsService) List() ([]model.Credentials, *Response, error) {
	root := new(model.Credentials)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

