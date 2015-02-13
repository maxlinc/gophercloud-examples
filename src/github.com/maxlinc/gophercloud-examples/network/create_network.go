package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/networks"
	"os"
)

func main() {
	ao := gophercloud.AuthOptions{
		IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		Username:         os.Getenv("RAX_USERNAME"),
		APIKey:           os.Getenv("RAX_API_KEY"),
	}
	provider, err := rackspace.AuthenticatedClient(ao)
	if err != nil {
		panic(err)
	}

	client, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RAX_REGION"),
	})
	if err != nil {
		panic(err)
	}

	// We specify a name and that it should forward packets
	opts := networks.CreateOpts{
		Label: "sample_network",
		CIDR:  "192.0.2.0/24",
	}

	// Execute the operation and get back a networks.Network struct
	network, err := networks.Create(client, opts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created network: %v (%v)", network.Label, network.CIDR)
}
