package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const organizationsBasePath = "organizations/"

type OrganizationsService struct {
	sling *sling.Sling
}

func NewOrganizationsService(sling *sling.Sling) *OrganizationsService {
	return &OrganizationsService {
		sling: sling.Path(organizationsBasePath),
	}
}

// List all organizations
func (s *OrganizationsService) List() ([]model.Organizations, *Response, error) {
	root := new(model.Organizations)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

