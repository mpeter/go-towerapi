package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const hostsBasePath = "hosts/"

type HostsService struct {
	sling *sling.Sling
}

func NewHostsService(sling *sling.Sling) *HostsService {
	return &HostsService {
		sling: sling.Path(hostsBasePath),
	}
}

// List all hosts
func (s *HostsService) List() ([]model.Hosts, *Response, error) {
	root := new(model.Hosts)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

