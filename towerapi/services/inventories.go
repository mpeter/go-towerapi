package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const inventoriesBasePath = "inventories/"

type InventoriesService struct {
	sling *sling.Sling
}

func NewInventoriesService(sling *sling.Sling) *InventoriesService {
	return &InventoriesService {
		sling: sling.Path(inventoriesBasePath),
	}
}

// List all inventories
func (s *InventoriesService) List() ([]model.Inventories, *Response, error) {
	root := new(model.Inventories)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

