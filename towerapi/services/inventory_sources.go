package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const inventorySourcesBasePath = "inventory_sources/"

type InventorySourcesService struct {
	sling *sling.Sling
}

func NewInventorySourcesService(sling *sling.Sling) *InventorySourcesService {
	return &InventorySourcesService {
		sling: sling.Path(inventorySourcesBasePath),
	}
}

// List all inventory_sources
func (s *InventorySourcesService) List() ([]model.InventorySources, *Response, error) {
	root := new(model.InventorySources)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

