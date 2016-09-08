package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const adHocCommandsBasePath = "ad_hoc_commands/"

type AdHocCommandsService struct {
	sling *sling.Sling
}

func NewAdHocCommandsService(sling *sling.Sling) *AdHocCommandsService {
	return &AdHocCommandsService {
		sling: sling.Path(adHocCommandsBasePath),
	}
}

// List all ad_hoc_commands
func (s *AdHocCommandsService) List() ([]model.AdHocCommands, *Response, error) {
	root := new(model.AdHocCommands)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

