package instance

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/utils/gnocchi"
	"github.com/gophercloud/utils/gnocchi/metric/v1/resources"
)

// Aggregate function returns aggregated resources
func Aggregate() {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := gnocchi.NewGnocchiV1(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	resources.List(client, resources.ListOpts{
		Limit:   10,
		SortKey: "id",
		SortDir: "desc",
	}, "instance").EachPage(func(page pagination.Page) (bool, error) {
		resourceList, err := resources.ExtractResources(page)
		if err != nil {
			fmt.Println(err)
			return false, err
		}

		for _, resource := range resourceList {
			fmt.Println(resource.OriginalResourceID)
			fmt.Println(resource.Metrics)
			fmt.Println(resource.EndedAt)
			fmt.Println(resource.StartedAt)
		}

		return true, nil
	})
}
