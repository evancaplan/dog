package dog

import (
	"context"
	"errors"
	"strconv"

	"github.com/digitalocean/godo"
)

type CreateDropletRequest struct {
	Name string
	Region
	DropletSize
	Image             string
	SSHKeys           []int
	Backups           bool
	IPv6              bool
	Configuration     string
	PrivateNetworking bool
	Volumes           []string
	Tags              []string
	VPCUUID           string
	Pat               string
}

type GetAllDropletsRequest struct {
	Page    int
	PerPage int
	Pat     string
}

type FindDropletByIDRequest struct {
	ID  int
	Pat string
}

type FindDropletsByTagRequest struct {
	Tag     string
	Page    int
	PerPage int
	Pat     string
}

type DeleteDropletRequest struct {
	ID int
	Pat string
}

// Droplet sizes
type DropletSize int

const (
	S1Cpu1GbRAM DropletSize = iota
	S1Cpu2GbRAM
	S1Cpu3GbRAM
	S2Cpu2GbRAM
	S3Cpu1GbRAM
	S2Cpu4GbRAM
	S4Cpu8GbRAM
	S6Cpu16GbRAM
	S8Cpu32GbRAM
	S12Cpu47GbRAM
	S16Cpu64GbRAM
	S20Cpu96GbRAM
	S24Cpu128GbRAM
	S32Cpu19GbRAM
)

func (ds DropletSize) String() string {
	names := [...]string{
		"s-1vcpu-1gb",
		"s-1vcpu-2gb",
		"s-1vcpu-3gb",
		"s-2vcpu-2gb",
		"s-3vcpu-1gb",
		"s-2vcpu-4gb",
		"s-4vcpu-8gb",
		"s-6vcpu-16gb",
		"s-8vcpu-32gb",
		"s-12vcpu-48gb",
		"s-16vcpu-64gb",
		"s-20vcpu-96gb",
		"s-24vcpu-128gb",
		"s-32vcpu-192gb",
	}
	if ds < S1Cpu1GbRAM || ds > S32Cpu19GbRAM {
		return "That is not a droplet size"
	}
	return names[ds]
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

func GetDropletById(fdr FindDropletByIDRequest) (*godo.Droplet, error) {

	c := Authenticate(fdr.Pat)
	ctx := context.TODO()

	droplet, _, err := c.Droplets.Get(ctx, fdr.ID)
	if err != nil {
		return nil, errors.New("Droplet with id: " + strconv.Itoa(fdr.ID) + ", was not found. Godo error: " + err.Error())
	}

	return droplet, nil
}

func GetDropletsByTag(fdr FindDropletsByTagRequest) (*[]godo.Droplet, error) {

	c := Authenticate(fdr.Pat)
	ctx := context.TODO()

	opt := &godo.ListOptions{
		Page:    fdr.Page,
		PerPage: fdr.PerPage,
	}

	droplets, _, err := c.Droplets.ListByTag(ctx, fdr.Tag, opt)
	if err != nil {
		return nil, errors.New("Droplets with Tag: " + fdr.Tag + ", weer not found. Godo error: " + err.Error())
	}

	return &droplets, nil
}

func CreateDroplet(cdr CreateDropletRequest) (*godo.Droplet, error) {

	keys := createGodoSSHKeys(cdr.SSHKeys)
	volumes := createVolumes(cdr.Volumes)

	create := &godo.DropletCreateRequest{
		Name:   cdr.Name,
		Region: cdr.Region.String(),
		Size:   cdr.DropletSize.String(),
		Image: godo.DropletCreateImage{
			Slug: cdr.Image,
		},
		SSHKeys:           keys,
		IPv6:              cdr.IPv6,
		UserData:          cdr.Configuration,
		PrivateNetworking: cdr.PrivateNetworking,
		Volumes:           volumes,
		Tags:              cdr.Tags,
		VPCUUID:           cdr.VPCUUID,
	}

	c := Authenticate(cdr.Pat)
	ctx := context.TODO()

	droplet, _, err := c.Droplets.Create(ctx, create)
	if err != nil {
		return nil, errors.New("Unable to create droplet. Godo error: " + err.Error())
	}

	return droplet, nil
}

func createGodoSSHKeys(keys []int) []godo.DropletCreateSSHKey {
	var godoKeys []godo.DropletCreateSSHKey

	for i := 1; i < len(keys); i++ {
		key := godo.DropletCreateSSHKey{ID: keys[i]}
		godoKeys = append(godoKeys, key)
	}
	return godoKeys
}

func createVolumes(volumes []string) []godo.DropletCreateVolume {
	var godoVolumes []godo.DropletCreateVolume

	for i := 1; i < len(volumes); i++ {
		volume := godo.DropletCreateVolume{ID: volumes[i]}
		godoVolumes = append(godoVolumes, volume)
	}
	return godoVolumes
}

func DeleteDroplet(ddr DeleteDropletRequest) error {
	
	c := Authenticate(ddr.Pat)
	ctx := context.TODO()

	_, err := c.Droplets.Delete(ctx,ddr.ID)
	if err != nil {
		return errors.New("Unable to delete droplet with ID: " + strconv.Itoa(ddr.ID))
	}
	return nil
}
