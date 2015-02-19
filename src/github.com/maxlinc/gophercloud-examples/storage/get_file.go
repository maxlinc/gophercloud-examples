package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/objectstorage/v1/objects"
	"io/ioutil"
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

	result := objects.Download(serviceClient, "{container_name}", "{remote_file}", nil)
	content, err := result.ExtractContent()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("somefile.txt", []byte(content), 0644)

	fmt.Printf("Retrieved file")
}
