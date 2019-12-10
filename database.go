package database

import (
	"context"
	"errors"

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

// Database Sizes
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

type GodoClient interface {
	Create(struct{})(*struct{}, error),
	GetById(id string, pat string)(*struct{}, error)
	GetAll(page int, numberPerPage int, pat string) (*[]struct{}, error)
}

type DatabaseClusterClient struct{
	Pat string
	client GodoClient
}

func(dbcc DatabaseClusterClient) Create(cdcr CreateDatabaseClusterRequest) (*godo.Datasbase, error) {
	
	// create new godo DatabaseCreateRequest
	create := &godo.DatabaseCreateRequest{
		Name:       cdcr.Name,
		EngineSlug: cdcr.DatabaseType.String(),
		Version:    cdcr.Version,
		SizeSlug:   cdcr.DatabaseSize.String(),
		NumNodes:   cdcr.NumNodes,
		Tags:       cdcr.Tags,
	}

	// generate new client and create empty context
	c := Authenticate(dbcc.Pat)
	ctx := context.TODO()

	// create new database cluster
	cluster, _, err := c.Databases.Create(ctx, create)
	if err != nil {
		return nil, errors.New("Unable to create database cluster. Godo error: " + err.Error())
	}

	return cluster, nil
}

func(dbcc DatabaseClusterClient) GetById(id string) (*godo.Database, error) {

	// generate new client and create empty context
	c := Authenticate(dbcc.Pat)
	ctx := context.TODO()

	// find database cluster by id
	cluster, _, err := c.Databases.Get(ctx, id)
	if err != nil {
		return nil, errors.New("Database cluster with id: " + id + " not found. Godo error: " + err.Error())
	}

	return cluster, nil
}

func(dbcc DatabaseClusterClient) GeAll(page int, numberPerPage int) (*[]godo.Database, error) {

	// create new godo ListOptions (page request)
	opt := &godo.ListOptions{
		Page:    page,
		PerPage: numberPerPage,
	}

	// generate new client and create empty context
	c := Authenticate(dbcc.Pat)
	ctx := context.TODO()

	// find all database clusters
	clusters, _, err := c.Databases.List(ctx, opt)
	if err != nil {
		return nil, errors.New("Unable to get all database clusters. Godo error: " + err.Error())
	}

	return &clusters, nil
}

func ResizeCluster(rcr ResizeClusterRequest) error {
	
	// create new godo ResizeDatabaseRequest
	resize := &godo.DatabaseResizeRequest{
		SizeSlug: rcr.DatabaseSize.String(),
		NumNodes: rcr.NumNodes,
	}

	// generate new client and create empty context
	c := Authenticate(rcr.Pat)
	ctx := context.TODO()

	// send resize request
	_, err := c.Databases.Resize(ctx, rcr.Id, resize)
	if err != nil {
		return errors.New("Unable to resize cluster" + rcr.Id + ". Godo error: " + err.Error())
	}

	return nil
}

func MigrateToNewRegion(mrr MigrateRegionRequest) error {

	// create new godo DatabaseMigrateRequest
	migrate := &godo.DatabaseMigrateRequest{
		Region: mrr.Region.String(),
	}
	
	// generate new client and create empty context
	c := Authenticate(mrr.Pat)
	ctx := context.TODO()

	// send migrate request
	_, err := c.Databases.Migrate(ctx, mrr.Id, migrate)
	if err != nil {
		return errors.New("Unable to migrate to a new region. Godo error: " + err.Error())
	}

	return nil
}

func ConfigureMaintenanceWindow(umw UpdateMaintenanceWindowRequest) error {

	// create new godo DatabaseUpdateMaintenanceRequest
	configure := &godo.DatabaseUpdateMaintenanceRequest{
		Day:  umw.Day,
		Hour: umw.Time,
	}
	
	// generate new client and create empty context
	c := Authenticate(umw.Pat)
	ctx := context.TODO()

	// send update maintanence window request
	_, err := c.Databases.UpdateMaintenance(ctx, umw.Id, configure)
	if err != nil {
		return errors.New("Unable to configure maintenance window for database cluster." + umw.Id + " Godo error: " + err.Error())
	}

	return nil
}

func addDatabaseToCluster(and AddNewDatabaseRequest) (*godo.DatabaseDB, error) {

	// create new godo DatabaseCreateDBRequest
	create := &godo.DatabaseCreateDBRequest{
		Name: and.Name,
	}

	// generate new client and create empty context
	c := Authenticate(and.Pat)
	ctx := context.TODO()

	// add database to cluster
	db, _, err := c.Databases.CreateDB(ctx, and.ClusterID, create)
	if err != nil {
		return nil, errors.New("Unable to add database to cluster. Godo error: " + err.Error())
	}

	return db, nil
}

func findAllDatabasesInCluster(pat string, clusterID string) (*[]godo.DatabaseDB, error) {
	// generate new client and create empty context
	c := Authenticate(pat)
	ctx := context.TODO()

	// find all databases by cluser id
	dbs, _, err := c.Databases.ListDBs(ctx, clusterID, nil)
	if err != nil {
		return nil, errors.New("Unable to find all databases in cluster: " + clusterID + " . Godo error:  " + err.Error())
	}

	return &dbs, nil
}

func deleteDatabaseInCluster(dr DeleteDatabaseRequest) error {
	
	// generate new client and create empty context
	c := Authenticate(dr.Pat)
	ctx := context.TODO()

	// send delete database request
	_, err := c.Databases.DeleteDB(ctx, dr.ClusterID, dr.Name)
	if err != nil {
		return errors.New("Unable to delete database: " + dr.Name + " . Godo error: " + err.Error())
	}

	return nil
}
