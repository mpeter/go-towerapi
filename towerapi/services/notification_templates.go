package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const notificationTemplatesBasePath = "notification_templates/"

type NotificationTemplatesService struct {
	sling *sling.Sling
}

func NewNotificationTemplatesService(sling *sling.Sling) *NotificationTemplatesService {
	return &NotificationTemplatesService {
		sling: sling.Path(notificationTemplatesBasePath),
	}
}

// List all notification_templates
func (s *NotificationTemplatesService) List() ([]model.NotificationTemplates, *Response, error) {
	root := new(model.NotificationTemplates)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

