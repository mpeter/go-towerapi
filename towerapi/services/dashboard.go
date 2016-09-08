package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const dashboardBasePath = "dashboard/"

type DashboardService struct {
	sling *sling.Sling
}

func NewDashboardService(sling *sling.Sling) *DashboardService {
	return &DashboardService {
		sling: sling.Path(dashboardBasePath),
	}
}

// List all dashboard
func (s *DashboardService) List() ([]model.Dashboard, *Response, error) {
	root := new(model.Dashboard)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

