package hosts

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/base"
	//"github.com/go-yaml/yaml"
	//	"strings"
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

// Request represents a request to create a new key.
type Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Inventory   int    `json:"inventory"`
	Enabled     bool   `json:"enabled"`
	InstanceID  string `json:"instance_id"`
	Variables   string `json:"variables" yaml:"variables"`
}

// Create creates a new Host
func (s *Service) Create(r *Request) (*Host, *http.Response, error) {
	host := new(Host)
	towerError := new(base.TowerError)
	response, err := s.sling.New().Post("").BodyJSON(r).Receive(host, towerError)

	return host, response, err
}

// Updates modifies an existing host
func (s *Service) Update(id int, r *Request) (*Host, *http.Response, error) {
	host := new(Host)
	towerError := new(base.TowerError)
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Put(path).BodyJSON(r).Receive(host, towerError)
	return host, response, err
}

func (s *Service) Delete(id int) (*http.Response, error) {
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Delete(path).ReceiveSuccess(nil)
	return response, err
}

func (s *Service) List() ([]Host, *http.Response, error) {
	hosts := new(Hosts)
	towerError := new(base.TowerError)
	response, err := s.sling.New().Get("").Receive(hosts, towerError)
	return hosts.Results, response, err
}

func (s *Service) ListByName(name string) ([]Host, *http.Response, error) {
	hosts := new(Hosts)
	towerError := new(base.TowerError)
	listParams := &base.ListParams{Name: name}
	response, err := s.sling.New().Get("").QueryStruct(listParams).Receive(hosts, towerError)
	return hosts.Results, response, err
}

func (s *Service) Exists(id int) (bool, *http.Response, error) {
	var exists bool = false
	inv, response, err := s.GetByID(id)
	if inv == nil {
		exists = true
	}
	return exists, response, err
}

func (s *Service) GetByName(name string) (*Host, *http.Response, error) {
	hosts := new(Hosts)
	towerError := new(base.TowerError)
	listParams := &base.ListParams{Name: name}
	response, err := s.sling.New().Get("").QueryStruct(listParams).Receive(hosts, towerError)
	if err != nil {
		return nil, response, fmt.Errorf("Unable to search by Name: %v", err)
	}
	if hosts.Count == 0 {
		return nil, response, nil
	} else if hosts.Count > 1 {
		return nil, response, fmt.Errorf("Exact search returned more than 1 result, shouldn't happen: %v", err)
	}
	results := &hosts.Results[0]
	return results, response, nil
}

func (s *Service) GetByID(id int) (*Host, *http.Response, error) {
	host := new(Host)
	towerError := new(base.TowerError)
	path := fmt.Sprintf("%d/", id)
	response, err := s.sling.New().Get(path).Receive(host, towerError)
	return host, response, err
}
