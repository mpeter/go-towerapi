package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const teamsBasePath = "teams/"

type TeamsService struct {
	sling *sling.Sling
}

func NewTeamsService(sling *sling.Sling) *TeamsService {
	return &TeamsService {
		sling: sling.Path(teamsBasePath),
	}
}

// List all teams
func (s *TeamsService) List() ([]model.Teams, *Response, error) {
	root := new(model.Teams)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

