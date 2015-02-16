package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/networks"
	"os"
)

func main() {
	// ao := gophercloud.AuthOptions{
	//  IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
	//  Username:         os.Getenv("RAX_USERNAME"),
	//  APIKey:           os.Getenv("RAX_API_KEY"),
	// }
	ao, err := rackspace.AuthOptionsFromEnv()
	provider, err := rackspace.AuthenticatedClient(ao)
	if err != nil {
		panic(err)
	}

	client, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RS_REGION"),
	})
	if err != nil {
		panic(err)
	}

	networks.List(client).EachPage(func(page pagination.Page) (bool, error) {
		networkList, err := networks.ExtractNetworks(page)
		if err != nil {
			panic(err)
		}

		for _, network := range networkList {
			fmt.Printf("Label: %s, CIDR: %s\n", network.Label, network.CIDR)
		}

		return true, nil
	})
}
