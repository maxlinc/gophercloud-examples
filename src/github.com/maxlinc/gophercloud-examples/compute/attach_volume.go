package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	osVolumes "github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumes"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/volumeattach"
	osServers "github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/blockstorage/v1/volumes"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
	"os"
)

func FindServer(client *gophercloud.ServiceClient, server_name string) (osServers.Server, error) {
	var server osServers.Server
	opts := osServers.ListOpts{Name: server_name}

	pager := servers.List(client, opts)

	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		for _, i := range serverList {
			server = i
			return false, nil
		}
		return true, nil
	})
	return server, err
}

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

	blockStorageClient, err := rackspace.NewBlockStorageV1(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RAX_REGION"),
	})
	if err != nil {
		panic(err)
	}

	computeClient, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RAX_REGION"),
	})
	if err != nil {
		panic(err)
	}

	opts := osVolumes.CreateOpts{
		Name: "photos",
		Size: 100,
	}
	vol, err := volumes.Create(blockStorageClient, opts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created volume %v\n", vol.Name)

	server, err := FindServer(computeClient, "{server_name}")
	if err != nil {
		panic(err)
	}

	createOpts := volumeattach.CreateOpts{
		VolumeID: vol.ID,
	}
	_, err = volumeattach.Create(computeClient, server.ID, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Attached %v to %v\n", vol.Name, server.Name)
}
