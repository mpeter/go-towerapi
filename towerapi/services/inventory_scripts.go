package towerapi


import (
	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/model"
)

const inventoryScriptsBasePath = "inventory_scripts/"

type InventoryScriptsService struct {
	sling *sling.Sling
}

func NewInventoryScriptsService(sling *sling.Sling) *InventoryScriptsService {
	return &InventoryScriptsService {
		sling: sling.Path(inventoryScriptsBasePath),
	}
}

// List all inventory_scripts
func (s *InventoryScriptsService) List() ([]model.InventoryScripts, *Response, error) {
	root := new(model.InventoryScripts)
	errorResponse := new(ErrorResponse)
	resp, err := s.sling.New().Get().Receive(root, errorResponse)
	return root.Results, resp, err
}

