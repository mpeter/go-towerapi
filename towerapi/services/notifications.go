package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const notificationsBasePath = "notifications/"

type NotificationsService struct {
	sling *sling.Sling
}

func NewNotificationsService(sling *sling.Sling) *NotificationsService {
	return &NotificationsService {
		sling: sling.Path(notificationsBasePath),
	}
}

// List all notifications
func (s *NotificationsService) List() ([]model.Notifications, *Response, error) {
	root := new(model.Notifications)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

