package inventories

import (
	"fmt"

	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/params"
	"github.com/mpeter/sling"
)

const basePath = "inventories/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth inventory related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.Path(basePath),
	}
}

// Create creates a new Inventory
func (s *Service) Create(r *Request) (*Inventory, error) {
	inventory := new(Inventory)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(inventory, apierr)
	return inventory, errors.BuildError(err, apierr)
}

// Updates modifies an existing inventory
func (s *Service) Update(r *Request) (*Inventory, error) {
	inventory := new(Inventory)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Put(r.ID+"/").BodyJSON(r).Receive(inventory, apierr)
	return inventory, errors.BuildError(err, apierr)
}

func (s *Service) Delete(id string) error {
	_, err := s.sling.New().Delete(id + "/").ReceiveSuccess(nil)
	return errors.BuildError(err, nil)
}

func (s *Service) List() ([]Inventory, error) {
	inventories := new(Inventories)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get("").Receive(inventories, apierr)
	return inventories.Results, errors.BuildError(err, apierr)
}

func (s *Service) ListByName(name string) ([]Inventory, error) {
	inventories := new(Inventories)
	apierr := new(errors.APIError)
	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(inventories, apierr)
	return inventories.Results, errors.BuildError(err, apierr)
}

func (s *Service) GetByName(name string) (*Inventory, error) {
	inventories := new(Inventories)
	apierr := new(errors.APIError)

	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(inventories, apierr)
	if err := errors.BuildError(err, apierr); err != nil {
		return nil, err
	}
	if inventories.Count == 0 {
		return nil, nil
	} else if inventories.Count > 1 {
		return nil, fmt.Errorf("Exact search returned more than 1 result for %s, should never happen", name)
	}
	results := inventories.Results[0]
	return &results, nil
}

func (s *Service) GetByID(id string) (*Inventory, error) {
	inventory := new(Inventory)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get(id+"/").Receive(inventory, apierr)
	return inventory, errors.BuildError(err, apierr)
}
