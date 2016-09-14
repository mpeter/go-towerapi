package groups

import (
	"fmt"

	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/params"
	"github.com/mpeter/sling"
)

const basePath = "groups/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth group related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.Path(basePath),
	}
}

// Create creates a new Group
func (s *Service) Create(r *Request) (*Group, error) {
	group := new(Group)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(group, apierr)
	return group, errors.BuildError(err, apierr)
}

// Updates modifies an existing group
func (s *Service) Update(r *Request) (*Group, error) {
	group := new(Group)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Put(r.ID+"/").BodyJSON(r).Receive(group, apierr)
	return group, errors.BuildError(err, apierr)
}

func (s *Service) Delete(id string) error {
	_, err := s.sling.New().Delete(id + "/").ReceiveSuccess(nil)
	return errors.BuildError(err, nil)
}

func (s *Service) List() ([]Group, error) {
	groups := new(Groups)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get("").Receive(groups, apierr)
	return groups.Results, errors.BuildError(err, apierr)
}

func (s *Service) ListByName(name string) ([]Group, error) {
	groups := new(Groups)
	apierr := new(errors.APIError)
	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(groups, apierr)
	return groups.Results, errors.BuildError(err, apierr)
}

func (s *Service) GetByName(name string) (*Group, error) {
	groups := new(Groups)
	apierr := new(errors.APIError)

	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(groups, apierr)
	if err := errors.BuildError(err, apierr); err != nil {
		return nil, err
	}
	if groups.Count == 0 {
		return nil, nil
	} else if groups.Count > 1 {
		return nil, fmt.Errorf("Exact search returned more than 1 result for %s, should never happen", name)
	}
	results := groups.Results[0]
	return &results, nil
}

func (s *Service) GetByID(id string) (*Group, error) {
	group := new(Group)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get(id+"/").Receive(group, apierr)
	return group, errors.BuildError(err, apierr)
}
