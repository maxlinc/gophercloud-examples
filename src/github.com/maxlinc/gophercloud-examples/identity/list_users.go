package main

import (
  "fmt"
  "os"
  "github.com/rackspace/gophercloud"
  "github.com/rackspace/gophercloud/pagination"
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
  client := rackspace.NewIdentityV2(provider)

  users.List(client).EachPage(func(page pagination.Page) (bool, error) {
    userList, err := osUsers.ExtractUsers(page)
    if err != nil {
      panic(err)
    }

    for _, user := range userList {
      fmt.Printf("ID: %s, Name: %s", user.ID, user.Username)
    }

    return true, nil
  })
}
