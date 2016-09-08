package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const meBasePath = "me/"

type MeService struct {
	sling *sling.Sling
}

func NewMeService(sling *sling.Sling) *MeService {
	return &MeService {
		sling: sling.Path(meBasePath),
	}
}

// List all me
func (s *MeService) List() ([]model.Me, *Response, error) {
	root := new(model.Me)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

