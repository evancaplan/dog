package dog_test

import (
	"context"
	"time"
	"testing"

	"github.com/digitalocean/godo"

	"gitlab.com/bwss/dog"
)

const TestPAT string = "TestPAT"

var ExpectedDB = godo.Database{
	ID:          "TestID-2131241",
	Name:        "dbtest",
	EngineSlug:  "mysql",
	VersionSlug: "11",
	Connection: &godo.DatabaseConnection{
		URI:      "Test URI",
		Database: "defaultdb",
		Host:     "Test Host",
		Port:     25060,
		User:     "test.user",
		Password: "TestPassword&",
		SSL:      true,
	},
	PrivateConnection: &godo.DatabaseConnection{
		URI:      "Test URI",
		Database: "defaultdb",
		Host:     "Test Host",
		Port:     25060,
		User:     "test.admin",
		Password: "&drowssaPtset",
		SSL:      true,
	},
	Users: []godo.DatabaseUser{
		godo.DatabaseUser{
			Name:     "doadmin",
			Role:     "primary",
			Password: "zt91mum075ofzyww",
		},
	},
	DBNames: []string{
		"defaultdb",
	},
	NumNodes:   3,
	RegionSlug: "sfo2",
	Status:     "online",
	CreatedAt:  time.Date(2019, 2, 26, 6, 12, 39, 0, time.UTC),
	MaintenanceWindow: &godo.DatabaseMaintenanceWindow{
		Day:         "monday",
		Hour:        "13:51:14",
		Pending:     false,
		Description: nil,
	},
	SizeSlug:           "test size slug",
	PrivateNetworkUUID: "test private network uuid",
	Tags:               []string{"production", "staging"},
}	

type MockGodoDatabaseSvc struct{}

func (m *MockGodoDatabaseSvc) get(context.Context, string) (*godo.Database, *godo.Response, error) {
	return &ExpectedDB, &godo.Response{}, nil
}

func TestGetById(t *testing.T) {
	expected := ExpectedDB
	returned := dog.
}

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
}
