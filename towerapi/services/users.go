package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const usersBasePath = "users/"

type UsersService struct {
	sling *sling.Sling
}

func NewUsersService(sling *sling.Sling) *UsersService {
	return &UsersService {
		sling: sling.Path(usersBasePath),
	}
}

// List all users
func (s *UsersService) List() ([]model.Users, *Response, error) {
	root := new(model.Users)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

