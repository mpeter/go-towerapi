package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const groupsBasePath = "groups/"

type GroupsService struct {
	sling *sling.Sling
}

func NewGroupsService(sling *sling.Sling) *GroupsService {
	return &GroupsService {
		sling: sling.Path(groupsBasePath),
	}
}

// List all groups
func (s *GroupsService) List() ([]model.Groups, *Response, error) {
	root := new(model.Groups)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

