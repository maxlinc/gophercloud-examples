package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	osusers "github.com/rackspace/gophercloud/openstack/identity/v2/users"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/identity/v2/users"
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
	client := rackspace.NewIdentityV2(provider)

	users.List(client).EachPage(func(page pagination.Page) (bool, error) {
		userList, err := osusers.Extractusers(page)
		if err != nil {
			return false, err
		}

		for _, user := range userList {
			if user.Username == "{user}" {
				userId := user.ID
				err := servers.Delete(serviceClient, server.ID)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Deleted user: %v (%v)", user.Username, userId)
				return false, nil
			}
		}

		return true, nil
	})
}
