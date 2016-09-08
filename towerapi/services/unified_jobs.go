package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const unifiedJobsBasePath = "unified_jobs/"

type UnifiedJobsService struct {
	sling *sling.Sling
}

func NewUnifiedJobsService(sling *sling.Sling) *UnifiedJobsService {
	return &UnifiedJobsService {
		sling: sling.Path(unifiedJobsBasePath),
	}
}

// List all unified_jobs
func (s *UnifiedJobsService) List() ([]model.UnifiedJobs, *Response, error) {
	root := new(model.UnifiedJobs)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

