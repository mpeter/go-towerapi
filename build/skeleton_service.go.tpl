package towerapi

import (
	"github.com/dghubble/sling"
)

const {{name_lcamel}}BasePath = "/{{name_uri}}"

type {{name_struct}}Service struct {
	sling *sling.Sling
}

// New{{name_struct}} creates a new instance of this service
func New{{name_struct}}Service(sling *sling.Sling) *{{name_struct}}Service {
	return &{{name_struct}}Service {
		sling: sling.Path({{name_lcamel}}BasePath),
	}
}

// List all {{name_uri}}
func (s *{{name_struct}}Service) List() ([]{{name_struct}}, *Response, error) {
	root := new({{name_struct}})
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get("/").Receive(root, errorResponse)
	return root.Results, resp, err
}

