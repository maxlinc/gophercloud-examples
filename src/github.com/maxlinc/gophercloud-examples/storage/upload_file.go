package main

import (
	"bufio"
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/objectstorage/v1/objects"
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

	f, err := os.Open("{local_file}")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	_, err = objects.Create(
		serviceClient,
		"{container_name}",
		"{remote_file}",
		reader,
		nil,
	).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uploaded file")
}
