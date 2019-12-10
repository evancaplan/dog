package database_test

import (
	"testing"

	"github.com/digitalocean/godo"

	"gitlab.com/bwss/dog"
)

const TestPAT string = "TestPAT"

func TestCreateDatabaseCluster(t *testing.T) {

	expectedDBCluster := &godo.Database{}
	testReq := dog.CreateDatabaseClusterRequest{
		Name:         "Test Request",
		DatabaseType: dog.MySQL,
		Version:      "Test Version",
		DatabaseSize: dog.DbS1Cpu1GbRAM10GbStorage,
		Region:       dog.FRA1,
		NumNodes:     1,
		Tags:         nil,
		Pat:          TestPAT,
	}

	returnedDBCluster, _ := dog.CreateDatabaseCluster(testReq)

	if &expectedDBCluster != &returnedDBCluster {
		t.Errorf(returnedDBCluster.Name)
	}
}
