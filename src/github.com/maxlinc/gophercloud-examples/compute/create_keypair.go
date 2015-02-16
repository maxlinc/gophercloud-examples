package main

import (
  "fmt"
  "os"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/compute/v2/keypairs"
  osKeyPairs "github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
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
  client, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
    Region: os.Getenv("RAX_REGION"),
  })
  if err != nil {
    panic(err)
  }

  opts := osKeyPairs.CreateOpts{
    Name: "foo",
  }

  keypair, err := keypairs.Create(client, opts).Extract()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Created %v\n", keypair.Name)
}
