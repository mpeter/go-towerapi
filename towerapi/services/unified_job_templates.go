package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const unifiedJobTemplatesBasePath = "unified_job_templates/"

type UnifiedJobTemplatesService struct {
	sling *sling.Sling
}

func NewUnifiedJobTemplatesService(sling *sling.Sling) *UnifiedJobTemplatesService {
	return &UnifiedJobTemplatesService {
		sling: sling.Path(unifiedJobTemplatesBasePath),
	}
}

// List all unified_job_templates
func (s *UnifiedJobTemplatesService) List() ([]model.UnifiedJobTemplates, *Response, error) {
	root := new(model.UnifiedJobTemplates)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

