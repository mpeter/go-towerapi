package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const pingBasePath = "ping/"

type PingService struct {
	sling *sling.Sling
}

func NewPingService(sling *sling.Sling) *PingService {
	return &PingService {
		sling: sling.Path(pingBasePath),
	}
}

// List all ping
func (s *PingService) List() ([]model.Ping, *Response, error) {
	root := new(model.Ping)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

