package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const configBasePath = "config/"

type ConfigService struct {
	sling *sling.Sling
}

func NewConfigService(sling *sling.Sling) *ConfigService {
	return &ConfigService {
		sling: sling.Path(configBasePath),
	}
}

// List all config
func (s *ConfigService) List() ([]model.Config, *Response, error) {
	root := new(model.Config)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

