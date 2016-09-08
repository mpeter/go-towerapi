package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const systemJobTemplatesBasePath = "system_job_templates/"

type SystemJobTemplatesService struct {
	sling *sling.Sling
}

func NewSystemJobTemplatesService(sling *sling.Sling) *SystemJobTemplatesService {
	return &SystemJobTemplatesService {
		sling: sling.Path(systemJobTemplatesBasePath),
	}
}

// List all system_job_templates
func (s *SystemJobTemplatesService) List() ([]model.SystemJobTemplates, *Response, error) {
	root := new(model.SystemJobTemplates)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

