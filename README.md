[![Build Status](https://travis-ci.org/mpeter/go-towerapi.svg)](https://travis-ci.org/mpeter/go-towerapi)

# TowerAPI

WORK IN PROGRESS

TowerAPI is a Go client library for accessing the Ansible Tower V1 API.

You can view the client API docs here: [http://godoc.org/github.com/mpeter/go-towerapi](http://godoc.org/github.com/mpeter/go-towerapi)

You can view Ansible Tower API docs here: [http://docs.ansible.com/ansible-tower/latest/html/towerapi/index.html](http://docs.ansible.com/ansible-tower/latest/html/towerapi/index.html)

## Usage

```go
import "github.com/mpeter/go-towerapi"
```

Create a new Ansible Tower client, then use the exposed services to
access different parts of the Ansible Tower API.

### Authentication

Currently, Basic Auth is the only method of
authenticating with the API. You can manage your users
inside the Ansible Tower Control GUI.

You can then use your token to create a new client:

```go
import (
  "net/http"
  "github.com/mpeter/go-towerapi/towerapi"
)

type Config struct {
  Endpoint string
  Username string
  Password string
}

func (c *Config) NewClient() (*towerapi.Client, error) {
  config := &towerapi.ClientConfig{
    Endpoint = c.Endpoint,
    Username = c.Username,
    Password = c.Password,
  }

  return towerapi.NewClient(http.DefaultClient, config)
}

```
## Examples

To create a new Organization:

```go
	config := new(towerapi.ClientConfig)
	config.Endpoint = c.Endpoint
	config.Password = c.Password
	config.Username = c.Username

	//return towerapi.NewClient(http.DefaultClient, config)

	api, err := towerapi.NewClient(http.DefaultClient, config)
	request := &organizations.Request{
		Name: "Org Name",
		Description: "Org Description",
	}
	org, resp, err := api.Organizations.Create(request)
	if err != nil {
		fmt.Errorf("ERROR: Organization.Create: %v", err)
	}
	fmt.Printf("Organization Struct: %v", org)
	fmt.Printf("Organization Response: %v", resp)
```

### Pagination

If a list of items is paginated by the API, you must request pages individually. For example, to fetch all Droplets:

TBD

## Versioning 

TBD

## Documentation
For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/mpeter/go-towerapi) documentation.


## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).
