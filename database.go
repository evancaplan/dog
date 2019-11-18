package dog

import (
	"context"

	"github.com/digitalocean/godo"
)

// Structs 

type CreateDatabaseClusterRequest struct {
	Name string
	DatabaseType
	Version string
	Region 
	NumNodes int
	Tags []string
	Pat string

}

// Enums

// Database types
type DatabaseType int

const(
	PostGres DatabaseType = iota
	Redis
	MySQL
)

func (dbt DatabaseType) String() string {
	names := [...]string {
		"pg",
		"redis",
		"mysql",
	}
	if dbt < PostGres || dbt > MySQL {
		return "That is not a database type"
	}
	return names[dbt]
}

// Droplet regions
type Region int

const (
	NYC1 Region = iota
	NYC2
	NYC3
	AMS2
	AMS3
	SFO1
	SFO2
	SGP1
	LON1
	FRA1
	TOR1
	BLR1
)

func(r Region) String() string {
	names := [...]string {
		"nyc1",
		"nyc2",
		"nyc3",
		"ams2",
		"ams3", 
		"sfo1",
		"sfo2",
		"sgp1",
		"lon1",
		"fra1",
		"tor1",
		"blr1",
	}
	if r < NYC1 || r > BLR1 {
		return "That is not a region"
	}
	return names[r]
}

// Droplet sizes
type DropletSize int 

const (
	S_1CPU_1GB_RAM Size = iota
	S_1CPU_2GB_RAM
	S_1CPU_3GB_RAM
	S_2CPU_2GB_RAM
	S_3CPU_1GB_RAM
	S_2CPU_4GB_RAM
	S_4CPU_8GB_RAM
	S_6CPU_16BG_RAM
	S_8CPU_32GB_RAM
	S_12CPU_47GB_RAM
	S_16CPU_64GB_RAM
	S_20CPU_96GB_RAM
	S_24CPU_128GB_RAM
	S_32CPU_19GB_RAM
)

func (ds DropletSize) String() string {
	names := [...]string {
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
}

func CreateDatabaseCluster(cdcr CreateDatabaseClusterRequest) *godo.Database{
	
	dcr := &godo.DatabaseCreateRequest {
		Name: cdcr.Name,
		EngineSlug: cdcr.DatabaseType.String(),
		Version: cdcr.Version,
		SizeSlug: "123",
		NumNodes: cdcr.NumNodes,
		Tags: cdcr.Tags,
	}

	c := Authenticate(cdcr.Pat)
	ctx := context.TODO()

	cluster, _, err := c.Databases.Create(ctx, dcr) 

	if err != nil {
		panic("aunt jemima")
	}

	return cluster
}


func GeAlltDatabaseClusters(page int, numberPerPage int, pat string) *[]godo.Database{
	
	// map user input of page and size to godo ListOptions object
	opt := &godo.ListOptions{
		Page:    page,
		PerPage: numberPerPage,
	}

	c := Authenticate(pat)
	ctx := context.TODO()

	// make call with client to GODO Databases
	clusters, _, err := c.Databases.List(ctx, opt)
	if err != nil {
		panic("yo")
	}
	return &clusters
}


