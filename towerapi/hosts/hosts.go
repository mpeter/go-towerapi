package hosts

import (
	"fmt"

	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/params"
	"github.com/mpeter/sling"
)

const basePath = "hosts/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth host related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.Path(basePath),
	}
}

// Create creates a new Host
func (s *Service) Create(r *Request) (*Host, error) {
	host := new(Host)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(host, apierr)
	return host, errors.BuildError(err, apierr)
}

// Updates modifies an existing host
func (s *Service) Update(r *Request) (*Host, error) {
	host := new(Host)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Put(r.ID+"/").BodyJSON(r).Receive(host, apierr)
	return host, errors.BuildError(err, apierr)
}

func (s *Service) Delete(id string) error {
	_, err := s.sling.New().Delete(id + "/").ReceiveSuccess(nil)
	return errors.BuildError(err, nil)
}

func (s *Service) List() ([]Host, error) {
	hosts := new(Hosts)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get("").Receive(hosts, apierr)
	return hosts.Results, errors.BuildError(err, apierr)
}

func (s *Service) ListByName(name string) ([]Host, error) {
	hosts := new(Hosts)
	apierr := new(errors.APIError)
	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(hosts, apierr)
	return hosts.Results, errors.BuildError(err, apierr)
}

func (s *Service) GetByName(name string) (*Host, error) {
	hosts := new(Hosts)
	apierr := new(errors.APIError)

	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(hosts, apierr)
	if err := errors.BuildError(err, apierr); err != nil {
		return nil, err
	}
	if hosts.Count == 0 {
		return nil, nil
	} else if hosts.Count > 1 {
		return nil, fmt.Errorf("Exact search returned more than 1 result for %s, should never happen", name)
	}
	results := &hosts.Results[0]
	return results, nil
}

func (s *Service) GetByID(id string) (*Host, error) {
	host := new(Host)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get(id+"/").Receive(host, apierr)
	return host, errors.BuildError(err, apierr)
}
