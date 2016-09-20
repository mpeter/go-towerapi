package projects

import (
	"fmt"

	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/params"
	"github.com/mpeter/sling"
)

const basePath = "projects/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth project related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.Path(basePath),
	}
}

// Create creates a new Project
func (s *Service) Create(r *Request) (*Project, error) {
	project := new(Project)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(project, apierr)
	return project, errors.BuildError(err, apierr)
}

// Updates modifies an existing project
func (s *Service) Update(r *Request) (*Project, error) {
	project := new(Project)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Put(r.ID+"/").BodyJSON(r).Receive(project, apierr)
	return project, errors.BuildError(err, apierr)
}

func (s *Service) Delete(id string) error {
	_, err := s.sling.New().Delete(id + "/").ReceiveSuccess(nil)
	return errors.BuildError(err, nil)
}

func (s *Service) List() ([]Project, error) {
	projects := new(Projects)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get("").Receive(projects, apierr)
	return projects.Results, errors.BuildError(err, apierr)
}

func (s *Service) ListByName(name string) ([]Project, error) {
	projects := new(Projects)
	apierr := new(errors.APIError)
	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(projects, apierr)
	return projects.Results, errors.BuildError(err, apierr)
}

func (s *Service) GetByName(name string) (*Project, error) {
	projects := new(Projects)
	apierr := new(errors.APIError)

	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(projects, apierr)
	if err := errors.BuildError(err, apierr); err != nil {
		return nil, err
	}
	if projects.Count == 0 {
		return nil, nil
	} else if projects.Count > 1 {
		return nil, fmt.Errorf("Exact search returned more than 1 result for %s, should never happen", name)
	}
	results := projects.Results[0]
	return &results, nil
}

func (s *Service) GetByID(id string) (*Project, error) {
	project := new(Project)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get(id+"/").Receive(project, apierr)
	return project, errors.BuildError(err, apierr)
}
