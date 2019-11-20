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
	DatabaseSize
	Region
	NumNodes int
	Tags     []string
	Pat      string
}

type ResizeClusterRequest struct {
	Id string
	DatabaseSize
	NumNodes int
	Pat      string
}

type MigrateRegionRequest struct {
	Id string
	Region
	Pat string
}

type UpdateMaintenanceWindowRequest struct {
	Id   string
	Day  string
	Time string
	Pat  string
}

type AddNewDatabaseRequest struct {
	Name      string
	ClusterID string
	Pat       string
}

type DeleteDatabaseRequest struct {
	Name      string
	ClusterID string
	Pat       string
}

// Enums

// Database types

type DatabaseType int

const (
	PostGres DatabaseType = iota
	Redis
	MySQL
)

func (dbt DatabaseType) String() string {
	names := [...]string{
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

func (r Region) String() string {
	names := [...]string{
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

// TODO remove this once Droplet is made
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

type DatabaseSize int

const (
	DbS1Cpu1GbRAM10GbStorage DatabaseSize = iota
	DbS1Cpu2GbRAM25GbStorage
	DbS2Cpu4GbRAM38GbStorage
	DbS4Cpu8GbRAM115GbStorage
	DbS6Cpu16GbRAM270GbStorage
	DbS8Cpu32GbRAM580GbStorage
	DbS16Cpu64GbRAM1120GbStorage
)

func (ds DatabaseSize) String() string {
	names := [...]string{
		"db-s-1vcpu-1gb",
		"db-s-1vcpu-2gb",
		"db-s-2vcpu-4gb",
		"db-s-4vcpu-8gb",
		"db-s-6vcpu-16gb",
		"db-s-8vcpu-32gb",
		"db-s-16vcpu-64gb",
	}
	if ds < DbS1Cpu1GbRAM10GbStorage || ds > DbS16Cpu64GbRAM1120GbStorage {
		return "That is not a database size"
	}
	return names[ds]
}

func CreateDatabaseCluster(cdcr CreateDatabaseClusterRequest) *godo.Database {

	create := &godo.DatabaseCreateRequest{
		Name:       cdcr.Name,
		EngineSlug: cdcr.DatabaseType.String(),
		Version:    cdcr.Version,
		SizeSlug:   cdcr.DatabaseSize.String(),
		NumNodes:   cdcr.NumNodes,
		Tags:       cdcr.Tags,
	}

	c := Authenticate(cdcr.Pat)
	ctx := context.TODO()

	cluster, _, err := c.Databases.Create(ctx, create)

	if err != nil {
		panic("aunt jemima")
	}

	return cluster
}

func GetDatabaseClusterById(id string, pat string) *godo.Database {

	c := Authenticate(pat)
	ctx := context.TODO()

	cluster, _, err := c.Databases.Get(ctx, id)

	if err != nil {

	}

	return cluster
}

func GeAllDatabaseClusters(page int, numberPerPage int, pat string) *[]godo.Database {

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

func ResizeCluster(rcr ResizeClusterRequest) {

	resize := &godo.DatabaseResizeRequest{
		SizeSlug: rcr.DatabaseSize.String(),
		NumNodes: rcr.NumNodes,
	}

	c := Authenticate(rcr.Pat)
	ctx := context.TODO()

	_, err := c.Databases.Resize(ctx, rcr.Id, resize)

	if err != nil {

	}
}

func MigrateToNewRegion(mrr MigrateRegionRequest) {

	migrate := &godo.DatabaseMigrateRequest{
		Region: mrr.Region.String(),
	}

	c := Authenticate(mrr.Pat)
	ctx := context.TODO()

	_, err := c.Databases.Migrate(ctx, mrr.Id, migrate)

	if err != nil {

	}
}

func ConfigureMaintenanceWindow(umw UpdateMaintenanceWindowRequest) {

	configure := &godo.DatabaseUpdateMaintenanceRequest{
		Day:  umw.Day,
		Hour: umw.Time,
	}

	c := Authenticate(umw.Pat)
	ctx := context.TODO()

	_, err := c.Databases.UpdateMaintenance(ctx, umw.Id, configure)

	if err != nil {

	}

}

func addDatabaseToCluster(and AddNewDatabaseRequest) *godo.DatabaseDB {

	create := &godo.DatabaseCreateDBRequest{
		Name: and.Name,
	}

	c := Authenticate(and.Pat)
	ctx := context.TODO()

	db, _, err := c.Databases.CreateDB(ctx, and.ClusterID, create)

	if err != nil {

	}
	return db
}

func findAllDatabasesInCluster(pat string, clusterID string) *[]godo.DatabaseDB {

	c := Authenticate(pat)
	ctx := context.TODO()

	dbs, _, err := c.Databases.ListDBs(ctx, clusterID, nil)
	if err != nil {

	}
	return &dbs
}

func deleteDatabaseInCluster(dr DeleteDatabaseRequest) {

	c := Authenticate(dr.Pat)
	ctx := context.TODO()

	_, err := c.Databases.DeleteDB(ctx, dr.ClusterID, dr.Name)

	if err != nil{

	}
}
