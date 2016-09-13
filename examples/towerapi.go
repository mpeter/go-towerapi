package main

import (
	"fmt"
	"github.com/mpeter/go-towerapi/towerapi"
	"github.com/mpeter/go-towerapi/towerapi/inventories"
	"log"
	"net/http"
	"os"
)

func main() {
	endpoint := os.Getenv("TOWER_ENDPOINT")
	username := os.Getenv("TOWER_USERNAME")
	password := os.Getenv("TOWER_PASSWORD")

	fmt.Println("Using these values:")
	fmt.Println("endpoint:", endpoint)
	fmt.Println("username:", username)
	fmt.Println("password:", password)

	if len(endpoint) == 0 || len(username) == 0 || len(password) == 0 {
		log.Fatal("Environment variable(s) not set\n")
	}
	client, err := towerapi.NewClient(http.DefaultClient, endpoint, username, password)
	if err != nil {
		log.Fatalf("towerapi.NewClient: %s", err)
	}
	anInventory := &inventories.Request{
		Name:         "aname15",
		Description:  "asdf",
		Organization: 1,
		Variables:    "",
	}
	inventory, result, err := client.Inventories.Create(anInventory)
	if err != nil {
		fmt.Errorf("Inventories.Create returned error: %v", err)
	}
	log.Printf("Inventories.Create Inventory: %v Response: %v", inventory, result)

	inventory1, _, err := client.Inventories.GetByID(inventory.ID)
	log.Printf("Inventory ID: %d", inventory1.ID)

	name := inventory1.Name
	inventory2, _, err := client.Inventories.GetByName(name)
	log.Printf("Inventory Name: %s", inventory2.Name)

	anInventory.Description = "changed description"
	inventory3, _, err := client.Inventories.Update(inventory2.ID, anInventory)
	log.Printf("Inventory Description: %s", inventory3.Description)

	client.Inventories.Delete(inventory3.ID)
	inventory4, result, err := client.Inventories.Create(anInventory)
	log.Printf("Inventory Name (After Delete): %s", inventory4.Name)

	// Cleanup
	client.Inventories.Delete(inventory4.ID)

}
