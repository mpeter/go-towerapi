package main

import (
	"fmt"
	"github.com/mpeter/go-towerapi/towerapi"
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
	client, _ := towerapi.NewClient(http.DefaultClient, endpoint, username, password)

	result, _, err := client.Organizations.List()
	if err != nil {
		fmt.Errorf("Organizations.List returned error: %v", err)
	}
	log.Printf("Organizations.List Response: %s", result)

	/*
			result1, _, err1 := client.Hosts.List(nil)
			if err1 != nil {
				fmt.Errorf("Hosts.List returned error: %v", err1)
			}
			log.Printf("Hosts.List Response: %s", result1)

			result2, _, err2 := client.Groups.List(nil)
			if err2 != nil {
				fmt.Errorf("Groups.List returned error: %v", err2)
			}
			log.Printf("Groups.List Response: %s", result2)
		result3, _, err3 := client.Inventories.List(nil)
		if err3 != nil {
			fmt.Errorf("Inventories.List returned error: %v", err3)
		}
		log.Printf("Inventories.List Response: %s", result3)
	*/

}
