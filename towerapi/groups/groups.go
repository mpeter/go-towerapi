package groups

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/base"
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

// Request represents a request to create a new key.
type Request struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Organization int    `json:"organization"`
	Variables    string `json:"variables" yaml:"variables"`
}

// Create creates a new Groups
func (s *Service) Create(r *Request) (*Groups, *http.Response, error) {
	group := new(Groups)
	towerError := new(base.TowerError)
	response, err := s.sling.New().Post("").BodyJSON(r).Receive(group, towerError)

	return group, response, err
}

// Updates modifies an existing group
func (s *Service) Update(id int, r *Request) (*Groups, *http.Response, error) {
	group := new(Groups)
	towerError := new(base.TowerError)
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Put(path).BodyJSON(r).Receive(group, towerError)
	return group, response, err
}

func (s *Service) Delete(id int) (*http.Response, error) {
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Delete(path).ReceiveSuccess(nil)
	return response, err
}

func (s *Service) List() ([]Groups, *http.Response, error) {
	groups := new(Groups)
	towerError := new(base.TowerError)
	response, err := s.sling.New().Get("").Receive(groups, towerError)
	return groups.Results, response, err
}

func (s *Service) ListByName(name string) ([]Groups, *http.Response, error) {
	groups := new(Groups)
	towerError := new(base.TowerError)
	listParams := &base.ListParams{Name: name}
	response, err := s.sling.New().Get("").QueryStruct(listParams).Receive(groups, towerError)
	return groups.Results, response, err
}

func (s *Service) Exists(id int) (bool, *http.Response, error) {
	var exists bool = false
	inv, response, err := s.GetByID(id)
	if inv == nil {
		exists = true
	}
	return exists, response, err
}

func (s *Service) GetByName(name string) (*Groups, *http.Response, error) {
	groups := new(Groups)
	towerError := new(base.TowerError)
	listParams := &base.ListParams{Name: name}
	response, err := s.sling.New().Get("").QueryStruct(listParams).Receive(groups, towerError)
	if err != nil {
		return nil, response, fmt.Errorf("Unable to search by Name: %v", err)
	}
	if groups.Count == 0 {
		return nil, response, nil
	} else if groups.Count > 1 {
		return nil, response, fmt.Errorf("Exact search returned more than 1 result, shouldn't happen: %v", err)
	}
	results := &groups.Results[0]
	return results, response, nil
}

func (s *Service) GetByID(id int) (*Groups, *http.Response, error) {
	group := new(Groups)
	towerError := new(base.TowerError)
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Get(path).Receive(group, towerError)
	return group, response, err
}
