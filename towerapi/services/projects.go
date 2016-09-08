package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const projectsBasePath = "projects/"

type ProjectsService struct {
	sling *sling.Sling
}

func NewProjectsService(sling *sling.Sling) *ProjectsService {
	return &ProjectsService {
		sling: sling.Path(projectsBasePath),
	}
}

// List all projects
func (s *ProjectsService) List() ([]model.Projects, *Response, error) {
	root := new(model.Projects)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

