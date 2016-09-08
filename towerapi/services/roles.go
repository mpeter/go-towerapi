package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const rolesBasePath = "roles/"

type RolesService struct {
	sling *sling.Sling
}

func NewRolesService(sling *sling.Sling) *RolesService {
	return &RolesService {
		sling: sling.Path(rolesBasePath),
	}
}

// List all roles
func (s *RolesService) List() ([]model.Roles, *Response, error) {
	root := new(model.Roles)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

