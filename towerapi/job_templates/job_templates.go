package job_templates

import (
	"fmt"

	"reflect"

	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/params"
	"github.com/mpeter/sling"
)

const basePath = "job_templates/"

// Service is an interface for interfacing with the
// endpoints of the Ansible Tower API
type Service struct {
	sling *sling.Sling
}

// NewService handles communication with auth job_template related methods of the
// Ansible Tower API.
func NewService(sling *sling.Sling) *Service {
	return &Service{
		sling: sling.Path(basePath),
	}
}

// Create creates a new JobTemplate
func create(s *sling.Sling, r interface{}, suc interface{}) (success interface{}, failure error) {
	result := reflect.New(reflect.TypeOf(success))
	apierr := new(errors.APIError)
	_, err := s.New().Post("").BodyJSON(r).Receive(result, apierr)
	return success, errors.BuildError(err, apierr).(error)
}

// Create creates a new JobTemplate
func (s *Service) Create(r *Request) (success *JobTemplate, e error) {
	//if res,err := create(s.sling, r, success) ; err != nil {
	//	return nil, err
	//} else {
	//	return res.(*JobTemplate), nil
	//}

	job_template := new(JobTemplate)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Post("").BodyJSON(r).Receive(job_template, apierr)
	return job_template, errors.BuildError(err, apierr)
}

// Updates modifies an existing job_template
func (s *Service) Update(r *Request) (*JobTemplate, error) {
	job_template := new(JobTemplate)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Put(r.ID+"/").BodyJSON(r).Receive(job_template, apierr)
	return job_template, errors.BuildError(err, apierr)
}

func (s *Service) Delete(id string) error {
	_, err := s.sling.New().Delete(id + "/").ReceiveSuccess(nil)
	return errors.BuildError(err, nil)
}

func (s *Service) List() ([]JobTemplate, error) {
	job_templates := new(JobTemplates)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get("").Receive(job_templates, apierr)
	return job_templates.Results, errors.BuildError(err, apierr)
}

func (s *Service) ListByName(name string) ([]JobTemplate, error) {
	job_templates := new(JobTemplates)
	apierr := new(errors.APIError)
	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(job_templates, apierr)
	return job_templates.Results, errors.BuildError(err, apierr)
}

func (s *Service) GetByName(name string) (*JobTemplate, error) {
	job_templates := new(JobTemplates)
	apierr := new(errors.APIError)

	params := &params.ListParams{Name: name}
	_, err := s.sling.New().Get("").QueryStruct(params).Receive(job_templates, apierr)
	if err := errors.BuildError(err, apierr); err != nil {
		return nil, err
	}
	if job_templates.Count == 0 {
		return nil, nil
	} else if job_templates.Count > 1 {
		return nil, fmt.Errorf("Exact search returned more than 1 result for %s, should never happen", name)
	}
	results := &job_templates.Results[0]
	return results, nil
}

func (s *Service) GetByID(id string) (*JobTemplate, error) {
	job_template := new(JobTemplate)
	apierr := new(errors.APIError)
	_, err := s.sling.New().Get(id+"/").Receive(job_template, apierr)
	return job_template, errors.BuildError(err, apierr)
}
