package organizations

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/base"
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

// Request represents a request to create a new key.
type Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create creates a new Organization
func (s *Service) Create(r *Request) (*Organization, *http.Response, error) {
	organization := new(Organization)
	towerError := new(base.TowerError)
	response, err := s.sling.New().Post("").BodyJSON(r).Receive(organization, towerError)

	return organization, response, err
}

// Updates modifies an existing organization
func (s *Service) Update(id int, r *Request) (*Organization, *http.Response, error) {
	organization := new(Organization)
	towerError := new(base.TowerError)
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Put(path).BodyJSON(r).Receive(organization, towerError)
	return organization, response, err
}

func (s *Service) Delete(id int) (*http.Response, error) {
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Delete(path).ReceiveSuccess(nil)
	return response, err
}

func (s *Service) List() ([]Organization, *http.Response, error) {
	organizations := new(Organizations)
	towerError := new(base.TowerError)
	response, err := s.sling.New().Get("").Receive(organizations, towerError)
	return organizations.Results, response, err
}

func (s *Service) ListByName(name string) ([]Organization, *http.Response, error) {
	organizations := new(Organizations)
	towerError := new(base.TowerError)
	listParams := &base.ListParams{Name: name}
	response, err := s.sling.New().Get("").QueryStruct(listParams).Receive(organizations, towerError)
	return organizations.Results, response, err
}

func (s *Service) Exists(id int) (bool, *http.Response, error) {
	var exists bool = false
	inv, response, err := s.GetByID(id)
	if inv == nil {
		exists = true
	}
	return exists, response, err
}

func (s *Service) GetByName(name string) (*Organization, *http.Response, error) {
	organizations := new(Organizations)
	towerError := new(base.TowerError)
	listParams := &base.ListParams{Name: name}
	response, err := s.sling.New().Get("").QueryStruct(listParams).Receive(organizations, towerError)
	if err != nil {
		return nil, response, fmt.Errorf("Unable to search by Name: %v", err)
	}
	if organizations.Count == 0 {
		return nil, response, nil
	} else if organizations.Count > 1 {
		return nil, response, fmt.Errorf("Exact search returned more than 1 result, shouldn't happen: %v", err)
	}
	results := &organizations.Results[0]
	return results, response, nil
}

func (s *Service) GetByID(id int) (*Organization, *http.Response, error) {
	organization := new(Organization)
	towerError := new(base.TowerError)
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Get(path).Receive(organization, towerError)
	return organization, response, err
}
