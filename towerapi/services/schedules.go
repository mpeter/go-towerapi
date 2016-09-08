package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const schedulesBasePath = "schedules/"

type SchedulesService struct {
	sling *sling.Sling
}

func NewSchedulesService(sling *sling.Sling) *SchedulesService {
	return &SchedulesService {
		sling: sling.Path(schedulesBasePath),
	}
}

// List all schedules
func (s *SchedulesService) List() ([]model.Schedules, *Response, error) {
	root := new(model.Schedules)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

