package main

import (
  "fmt"
  "os"
  "github.com/rackspace/gophercloud"
  osUsers "github.com/rackspace/gophercloud/openstack/identity/v2/users"
  "github.com/rackspace/gophercloud/rackspace"
  "github.com/rackspace/gophercloud/rackspace/identity/v2/users"
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
  // Should return err according to devsite, but only has one return value
  client := rackspace.NewIdentityV2(provider)

  opts := users.CreateOpts{
    Username: "{user}",
    Email: "{email}",
    Enabled: osUsers.Enabled,
  }

  user, err := users.Create(client, opts).Extract()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Created %v\n", user.Username)
}
