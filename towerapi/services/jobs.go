package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const jobsBasePath = "jobs/"

type JobsService struct {
	sling *sling.Sling
}

func NewJobsService(sling *sling.Sling) *JobsService {
	return &JobsService {
		sling: sling.Path(jobsBasePath),
	}
}

// List all jobs
func (s *JobsService) List() ([]model.Jobs, *Response, error) {
	root := new(model.Jobs)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

