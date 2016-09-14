package organizations

import (
	"fmt"

	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/params"
	"github.com/mpeter/sling"
)

const basePath = "organizations/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth organization related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.Path(basePath),
	}
}

// Create creates a new Organization
func (s *Service) Create(r *Request) (*Organization, error) {
	organization := new(Organization)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(organization, apierr)
	return organization, errors.BuildError(err, apierr)
}

// Updates modifies an existing organization
func (s *Service) Update(r *Request) (*Organization, error) {
	organization := new(Organization)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Put(r.ID+"/").BodyJSON(r).Receive(organization, apierr)
	return organization, errors.BuildError(err, apierr)
}

func (s *Service) Delete(id string) error {
	_, err := s.sling.New().Delete(id + "/").ReceiveSuccess(nil)
	return errors.BuildError(err, nil)
}

func (s *Service) List() ([]Organization, error) {
	organizations := new(Organizations)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get("").Receive(organizations, apierr)
	return organizations.Results, errors.BuildError(err, apierr)
}

func (s *Service) ListByName(name string) ([]Organization, error) {
	organizations := new(Organizations)
	apierr := new(errors.APIError)
	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(organizations, apierr)
	return organizations.Results, errors.BuildError(err, apierr)
}

func (s *Service) GetByName(name string) (*Organization, error) {
	organizations := new(Organizations)
	apierr := new(errors.APIError)

	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(organizations, apierr)
	if err := errors.BuildError(err, apierr); err != nil {
		return nil, err
	}
	if organizations.Count == 0 {
		return nil, nil
	} else if organizations.Count > 1 {
		return nil, fmt.Errorf("Exact search returned more than 1 result for %s, should never happen", name)
	}
	results := organizations.Results[0]
	return &results, nil
}

func (s *Service) GetByID(id string) (*Organization, error) {
	organization := new(Organization)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get(id+"/").Receive(organization, apierr)
	return organization, errors.BuildError(err, apierr)
}
