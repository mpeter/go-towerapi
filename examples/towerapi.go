package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mpeter/go-towerapi/towerapi"
	"github.com/mpeter/go-towerapi/towerapi/inventories"
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
	config := &towerapi.Config{
		Endpoint: endpoint,
		Username: username,
		Password: password,
	}
	client, err := towerapi.NewClient(config)

	if err != nil {
		log.Fatalf("towerapi.NewClient: %s", err)
	}
	org := 1
	anInventory := &inventories.Request{
		Name:         "aname15",
		Description:  "asdf",
		Organization: &org,
		Variables:    "",
	}
	inventory, err := client.Inventories.Create(anInventory)
	if err != nil {
		fmt.Errorf("Inventories.Create returned error: %v", err)
	}
	log.Printf("Inventories.Create Inventory: %v", inventory)

	inventory1, err := client.Inventories.GetByID(strconv.Itoa(inventory.ID))
	log.Printf("Inventory ID: %d", inventory1.ID)

	name := inventory1.Name
	inventory2, err := client.Inventories.GetByName(name)
	log.Printf("Inventory Name: %s", inventory2.Name)

	anInventory.Description = "changed description"
	inventory3, err := client.Inventories.Update(anInventory)
	log.Printf("Inventory Description: %s", inventory3.Description)

	client.Inventories.Delete(strconv.Itoa(inventory3.ID))
	inventory4, err := client.Inventories.Create(anInventory)
	log.Printf("Inventory Name (After Delete): %s", inventory4.Name)

	// Cleanup
	client.Inventories.Delete(strconv.Itoa(inventory4.ID))

}
