package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const activityStreamBasePath = "activity_stream/"

type ActivityStreamService struct {
	sling *sling.Sling
}

func NewActivityStreamService(sling *sling.Sling) *ActivityStreamService {
	return &ActivityStreamService {
		sling: sling.Path(activityStreamBasePath),
	}
}

// List all activity_stream
func (s *ActivityStreamService) List() ([]model.ActivityStream, *Response, error) {
	root := new(model.ActivityStream)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

