package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/objectstorage/v1/containers"
	"os"
)

func main() {
	ao := gophercloud.AuthOptions{
		IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		Username:         os.Getenv("RAX_USERNAME"),
		APIKey:           os.Getenv("RAX_API_KEY"),
	}

	provider, err := rackspace.AuthenticatedClient(ao)

	serviceClient, err := rackspace.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RAX_REGION"),
	})

	_, err = containers.Create(serviceClient, "{container_name}", nil).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created container {container_name}")
}
