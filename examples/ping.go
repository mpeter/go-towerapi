package main

import (
	"fmt"
	"github.com/mpeter/go-towerapi/towerapi"
	"log"
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

	client, _ := towerapi.NewClient(endpoint, username, password)

	ping, _, err := client.Ping.Get()
	if err != nil {
		fmt.Errorf("Ping.Get returned error: %v", err)
	}

	log.Printf("Ping.Get Response: %s", ping)
}
