package dog

import (
	"context"
	"errors"
	"strconv"

	"github.com/digitalocean/godo"
)

type GetAllDropletsRequest struct {
	Page    int
	PerPage int
	Pat     string
}

type FindDropletRequest struct {
	ID  int
	Pat string
}

func GetAllDroplets(gar GetAllDropletsRequest) (*[]godo.Droplet, error) {

	opt := &godo.ListOptions{
		Page:    gar.Page,
		PerPage: gar.PerPage,
	}

	c := Authenticate(gar.Pat)
	ctx := context.TODO()

	droplets, _, err := c.Droplets.List(ctx, opt)
	if err != nil {
		return nil, errors.New("Unable to get all databases. Godo error: " + err.Error())
	}

	return &droplets, nil
}

func GetDropletById(fdr FindDropletRequest) (*godo.Droplet, error) {

	c := Authenticate(fdr.Pat)
	ctx := context.TODO()

	droplet, _, err := c.Droplets.Get(ctx, fdr.ID)
	if err != nil {
		return nil, errors.New("Droplet with id: " + strconv.Itoa(fdr.ID) + ", was not found. Godo error: " + err.Error())
	}

	return droplet, nil
}
