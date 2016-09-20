package job_templates

import (
	"reflect"
	"testing"

	"github.com/mpeter/sling"
)

func TestNewService(t *testing.T) {
	type args struct {
		sling *sling.Sling
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewService(tt.args.sling); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewService() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestService_Create(t *testing.T) {
	type fields struct {
		sling *sling.Sling
	}
	type args struct {
		r *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Service{
			sling: tt.fields.sling,
		}
		got, err := s.Create(tt.args.r)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.Create() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Service.Create() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestService_Update(t *testing.T) {
	type fields struct {
		sling *sling.Sling
	}
	type args struct {
		r *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Service{
			sling: tt.fields.sling,
		}
		got, err := s.Update(tt.args.r)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.Update() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Service.Update() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestService_Delete(t *testing.T) {
	type fields struct {
		sling *sling.Sling
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Service{
			sling: tt.fields.sling,
		}
		if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.Delete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestService_List(t *testing.T) {
	type fields struct {
		sling *sling.Sling
	}
	tests := []struct {
		name    string
		fields  fields
		want    []JobTemplate
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Service{
			sling: tt.fields.sling,
		}
		got, err := s.List()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.List() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Service.List() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestService_ListByName(t *testing.T) {
	type fields struct {
		sling *sling.Sling
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []JobTemplate
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Service{
			sling: tt.fields.sling,
		}
		got, err := s.ListByName(tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.ListByName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Service.ListByName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestService_GetByName(t *testing.T) {
	type fields struct {
		sling *sling.Sling
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Service{
			sling: tt.fields.sling,
		}
		got, err := s.GetByName(tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.GetByName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Service.GetByName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestService_GetByID(t *testing.T) {
	type fields struct {
		sling *sling.Sling
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *JobTemplate
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := &Service{
			sling: tt.fields.sling,
		}
		got, err := s.GetByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Service.GetByID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Service.GetByID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
