package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/sling"
	"golang.org/x/oauth2"

	"github.com/mpeter/go-towerapi/towerapi"
)

func main() {

	// Github OAuth2 API
	flags := flag.NewFlagSet("tower-auth", flag.ExitOnError)
	// -endpoint=xxx or TOWER_ENDPOINT env var
	endpoint := flags.String("endpoint", "", "Tower API endpoint")
	username := flags.String("username", "", "Tower API username")
	password := flags.String("password", "", "Tower API password")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TOWER")

	client, _ := towerapi.NewClient(endpoint, username, password)
	ping, _, err := client.Ping.Get()
	fmt.Printf("Your Github Issues:\n%v\n", ping)

}
