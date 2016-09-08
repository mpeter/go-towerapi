package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const systemJobsBasePath = "system_jobs/"

type SystemJobsService struct {
	sling *sling.Sling
}

func NewSystemJobsService(sling *sling.Sling) *SystemJobsService {
	return &SystemJobsService {
		sling: sling.Path(systemJobsBasePath),
	}
}

// List all system_jobs
func (s *SystemJobsService) List() ([]model.SystemJobs, *Response, error) {
	root := new(model.SystemJobs)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

