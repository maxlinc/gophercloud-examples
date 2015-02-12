package main

import (
	"fmt"
  "os"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
)

func main() {
	ao := gophercloud.AuthOptions{
    IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		Username: os.Getenv("RAX_USERNAME"),
		APIKey:   os.Getenv("RAX_API_KEY"),
	}
	provider, err := rackspace.AuthenticatedClient(ao)
	if err != nil {
		panic(err)
	}
  rackspace.NewIdentityV2(provider)

	_, err = rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RAX_REGION"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Authenticated.")
}
