package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const jobEventsBasePath = "job_events/"

type JobEventsService struct {
	sling *sling.Sling
}

func NewJobEventsService(sling *sling.Sling) *JobEventsService {
	return &JobEventsService {
		sling: sling.Path(jobEventsBasePath),
	}
}

// List all job_events
func (s *JobEventsService) List() ([]model.JobEvents, *Response, error) {
	root := new(model.JobEvents)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

