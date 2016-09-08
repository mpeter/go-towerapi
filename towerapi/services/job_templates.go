package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const jobTemplatesBasePath = "job_templates/"

type JobTemplatesService struct {
	sling *sling.Sling
}

func NewJobTemplatesService(sling *sling.Sling) *JobTemplatesService {
	return &JobTemplatesService {
		sling: sling.Path(jobTemplatesBasePath),
	}
}

// List all job_templates
func (s *JobTemplatesService) List() ([]model.JobTemplates, *Response, error) {
	root := new(model.JobTemplates)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

