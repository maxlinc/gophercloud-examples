package main

import (
  "fmt"
  "os"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/openstack/networking/v2/networks"
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

  client, err := rackspace.NewNetworkV2(provider, gophercloud.EndpointOpts{
    Name:   "neutron",
    Region: os.Getenv("RAX_REGION"),
  })

  // We specify a name and that it should forward packets
  opts := networks.CreateOpts{Name: "main_network", AdminStateUp: networks.Up}

  // Execute the operation and get back a networks.Network struct
  network, err := networks.Create(client, opts).Extract()
  fmt.Printf("Created network: %v (%v)", network.Name, network.ID)
}
