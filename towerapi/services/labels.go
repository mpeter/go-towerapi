package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const labelsBasePath = "labels/"

type LabelsService struct {
	sling *sling.Sling
}

func NewLabelsService(sling *sling.Sling) *LabelsService {
	return &LabelsService {
		sling: sling.Path(labelsBasePath),
	}
}

// List all labels
func (s *LabelsService) List() ([]model.Labels, *Response, error) {
	root := new(model.Labels)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

