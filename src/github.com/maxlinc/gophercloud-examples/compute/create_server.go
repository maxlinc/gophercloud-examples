package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	osFlavors "github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
	osImages "github.com/rackspace/gophercloud/openstack/compute/v2/images"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/flavors"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/images"
	"github.com/rackspace/gophercloud/rackspace/compute/v2/servers"
	"os"
)

func FindImage(client *gophercloud.ServiceClient, image_name string) osImages.Image {
	var image osImages.Image
	opts := osImages.ListOpts{Name: image_name}

	pager := images.ListDetail(client, opts)

	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := images.ExtractImages(page)
		if err != nil {
			return false, err
		}

		for _, i := range imageList {
			image = i
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		panic(err)
	}

	return image
}

func FindFlavor(client *gophercloud.ServiceClient, flavor_name string) osFlavors.Flavor {
	var flavor osFlavors.Flavor
	opts := osFlavors.ListOpts{}
	pager := flavors.ListDetail(client, opts)

	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		flavorList, err := flavors.ExtractFlavors(page)
		if err != nil {
			return false, err
		}

		for _, i := range flavorList {
			if i.Name == flavor_name {
				flavor = i
				return false, nil
			}
		}
		return true, nil
	})
	if err != nil {
		panic(err)
	}

	return flavor
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
	client, err := rackspace.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("RAX_REGION"),
	})
	if err != nil {
		panic(err)
	}

	flavor := FindFlavor(client, "{server_flavor}")
	image := FindImage(client, "{server_image}")

	server, err := servers.Create(client, servers.CreateOpts{
		Name:      "{server_name}",
		ImageRef:  image.ID,
		FlavorRef: flavor.ID,
	}).Extract()
	if err != nil {
		panic(err)
	}
	// Docs say this returns err, but it doesn't
	servers.WaitForStatus(client, server.ID, "ACTIVE", 600)

	fmt.Printf("Created %v (%v)\n", server.Name, server.ID)
}
