package credentials

import (
	"fmt"

	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/params"
	"github.com/mpeter/sling"
)

const basePath = "credentials/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth credential related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.Path(basePath),
	}
}

// Create creates a new Credential
func (s *Service) Create(r *Request) (*Credential, error) {
	credential := new(Credential)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(credential, apierr)
	return credential, errors.BuildError(err, apierr)
}

// Updates modifies an existing credential
func (s *Service) Update(r *Request) (*Credential, error) {
	credential := new(Credential)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Put(r.ID+"/").BodyJSON(r).Receive(credential, apierr)
	return credential, errors.BuildError(err, apierr)
}

func (s *Service) Delete(id string) error {
	_, err := s.sling.New().Delete(id + "/").ReceiveSuccess(nil)
	return errors.BuildError(err, nil)
}

func (s *Service) List() ([]Credential, error) {
	credentials := new(Credentials)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get("").Receive(credentials, apierr)
	return credentials.Results, errors.BuildError(err, apierr)
}

func (s *Service) ListByName(name string) ([]Credential, error) {
	credentials := new(Credentials)
	apierr := new(errors.APIError)
	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(credentials, apierr)
	return credentials.Results, errors.BuildError(err, apierr)
}

func (s *Service) GetByName(name string) (*Credential, error) {
	credentials := new(Credentials)
	apierr := new(errors.APIError)

	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(credentials, apierr)
	if err := errors.BuildError(err, apierr); err != nil {
		return nil, err
	}
	if credentials.Count == 0 {
		return nil, nil
	} else if credentials.Count > 1 {
		return nil, fmt.Errorf("Exact search returned more than 1 result for %s, should never happen", name)
	}
	results := credentials.Results[0]
	return &results, nil
}

func (s *Service) GetByID(id string) (*Credential, error) {
	credential := new(Credential)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get(id+"/").Receive(credential, apierr)
	return credential, errors.BuildError(err, apierr)
}
