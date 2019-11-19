package dog

import (
	"context"

	"github.com/digitalocean/godo"
)

type GetAllDropletsRequest struct {
	Page    int
	PerPage int
	Pat     string
}
 
func GetAllDroplets(gar GetAllDropletsRequest) *[]godo.Droplet {

	opt := &godo.ListOptions{
		Page:    gar.Page,
		PerPage: gar.PerPage,
	}

	c := Authenticate(gar.Pat)
	ctx := context.TODO()

	droplets, _, err := c.Droplets.List(ctx, opt)
	if err != nil {

	}

	return &droplets
}

