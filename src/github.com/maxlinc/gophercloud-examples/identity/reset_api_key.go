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
      return false, err
    }

    for _, user := range userList {
      if user.Username == "{user}" {
        userId := user.ID
        newKey, err := users.ResetAPIKey(client, userId).Extract()
        if err != nil {
          panic(err)
        }
        fmt.Printf("New API Key: %v", newKey)
        return false, nil
      }
    }

    return true, nil
  })
}
