package search

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/utils/gnocchi/metric/v1/resources"
)

const searchPath = "search"
const resourcePath = "resource"

// Search returns a Pager
func Search(c *gophercloud.ServiceClient, opts resources.ListOptsBuilder, resourceType string) pagination.Pager {
	url := c.ServiceURL(searchPath, resourcePath, resourceType)
	fmt.Println(url)

	return pagination.Pager{}
}
