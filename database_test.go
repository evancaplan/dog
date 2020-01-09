package dog

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/digitalocean/godo"
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

var ExpectedDBs = []godo.Database{ExpectedDB}

var ExpectedDatabaseDB = godo.DatabaseDB{
	Name: "Test Database Name",
}

var ExpectedDatabaseDBs = []godo.DatabaseDB{ExpectedDatabaseDB}

var TestError = "test error"

var TestCreateDatabeClusterRequest = CreateDatabaseClusterRequest{
	Name:         "Test Request",
	DatabaseType: MySQL,
	Version:      "Test Version",
	DatabaseSize: DbS1Cpu1GbRAM10GbStorage,
	Region:       FRA1,
	NumNodes:     1,
	Tags:         nil,
}

var TestResizeClusterRequest = ResizeClusterRequest{
	Id:           "1",
	DatabaseSize: DbS1Cpu2GbRAM25GbStorage,
	NumNodes:     1,
}

var TestMigrateNewRegionRequest = MigrateRegionRequest{
	Id:     "1",
	Region: NYC2,
}

var TestUpdateMaintenanceWindowRequest = UpdateMaintenanceWindowRequest{
	Id:   "1",
	Day:  "Tuesday",
	Time: "18:00",
}

var TestCreateDatabaseRequest = CreateDatabaseRequest{
	Name:      "Test Database Name",
	ClusterID: "234234-1234BC24",
}

var TestDeleteDatabaseRequest = DeleteDatabaseRequest{
	Name:      "Test Database Name",
	ClusterID: "123-4567",
}

func TestGetById(t *testing.T) {

	dbClient := NewDBC(TestPAT)
	dbClient.client = &MockGodoDatabaseSvc{}
	
	expected := ExpectedDB
	returned, _ := dbClient.GetById("1")
	if !reflect.DeepEqual(&expected, returned) {
		t.Errorf("expected %+v\n returned %+v\n", expected, returned)
	}

}

func TestGetAll(t *testing.T) {

	dbClient := NewDBC(TestPAT)
	dbClient.client = &MockGodoDatabaseSvc{}
	
	expected := ExpectedDBs
	returned, _ := dbClient.GetAll(1, 5)
	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("expected %+v\n returned %+v\n", expected, returned)
	}

}

func TestCreateDatabaseCluster(t *testing.T) {

	dbClient := NewDBC(TestPAT)
	dbClient.client = &MockGodoDatabaseSvc{}
	expected := ExpectedDB

	returned, _ := dbClient.Create(TestCreateDatabeClusterRequest)
	if !reflect.DeepEqual(&expected, returned) {
		t.Errorf("expected %+v\n returned %+v\n", expected, returned)
	}

}

func TestResizeCluster(t *testing.T) {

	t.Run("Error is thrown", func(t *testing.T) {
		dbClient := NewDBC(TestPAT)
		dbClient.client = &MockGodoDatabaseSvc{}

		expectedError := "Unable to resize cluster " + TestResizeClusterRequest.Id + ". Godo error: " + TestError
		returnedError := dbClient.ResizeCluster(TestResizeClusterRequest).Error()
		if returnedError != expectedError {
			t.Errorf("expected: %s returned: %s", expectedError, returnedError)
		}
	})

}

func TestMigrateToNewRegion(t *testing.T) {

	t.Run("Error is thrown", func(t *testing.T) {
		dbClient := NewDBC(TestPAT)
		dbClient.client = &MockGodoDatabaseSvc{}

		expectedError := "Unable to migrate to new region. Godo error: test error"
		returnedError := dbClient.MigrateToNewRegion(TestMigrateNewRegionRequest).Error()
		if returnedError != expectedError {
			t.Errorf("expected: %s returned: %s", expectedError, returnedError)
		}
	})

}

func TestConfigureMaintenanceWindow(t *testing.T) {

	t.Run("Error is thrown", func(t *testing.T) {
		dbClient := NewDBC(TestPAT)
		dbClient.client = &MockGodoDatabaseSvc{}

		expectedError := "Unable to configure maintenance window for database cluster." + TestUpdateMaintenanceWindowRequest.Id + " Godo error: " + TestError
		returnedError := dbClient.ConfigureMaintenanceWindow(TestUpdateMaintenanceWindowRequest).Error()
		if expectedError != returnedError {
			t.Errorf("expected: %s returned: %s", expectedError, returnedError)
		}
	})

}

func TestAddDatabaseToCluster(t *testing.T) {

	dbClient := NewDBC(TestPAT)
	dbClient.client = &MockGodoDatabaseSvc{}

	expected := ExpectedDatabaseDB
	returned, _ := dbClient.AddDatabaseToCluster(TestCreateDatabaseRequest)
	if !reflect.DeepEqual(&expected, returned) {
		t.Errorf("expected: %+v\n returned: %+v\n", expected, returned)
	}

}

func TestFindAllDatabasesInCluster(t *testing.T) {

	dbClient := NewDBC(TestPAT)
	dbClient.client = &MockGodoDatabaseSvc{}

	expected := ExpectedDatabaseDBs
	returned, _ := dbClient.FindAllDatabasesInCluster("1")
	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("expected: %+v\n returned: %+v\n", expected, returned)
	}

}

func TestDeleteDatabaseInCluster(t *testing.T) {

	t.Run("Error is thrown", func(t *testing.T) {
		dbClient := NewDBC(TestPAT)
		dbClient.client = &MockGodoDatabaseSvc{}

		expectedError := "Unable to delete database: " + TestDeleteDatabaseRequest.Name + " . Godo error: test error"
		returnedError := dbClient.DeleteDatabaseInCluster(TestDeleteDatabaseRequest).Error()
		if expectedError != returnedError {
			t.Errorf("expected: %s\n returned: %s\n", expectedError, returnedError)
		}
	})

}

type MockGodoDatabaseSvc struct{}

func (m *MockGodoDatabaseSvc) Get(context.Context, string) (*godo.Database, *godo.Response, error) {
	return &ExpectedDB, nil, nil
}

func (m *MockGodoDatabaseSvc) Create(context.Context, *godo.DatabaseCreateRequest) (*godo.Database, *godo.Response, error) {
	return &ExpectedDB, nil, nil
}

func (m *MockGodoDatabaseSvc) CreateDB(context.Context, string, *godo.DatabaseCreateDBRequest) (*godo.DatabaseDB, *godo.Response, error) {
	return &ExpectedDatabaseDB, nil, nil
}

func (m *MockGodoDatabaseSvc) DeleteDB(context.Context, string, string) (*godo.Response, error) {
	return nil, errors.New(TestError)
}

func (m *MockGodoDatabaseSvc) GetDB(context.Context, string, string) (*godo.DatabaseDB, *godo.Response, error) {
	return nil, nil, nil
}

func (m *MockGodoDatabaseSvc) List(context.Context, *godo.ListOptions) ([]godo.Database, *godo.Response, error) {
	return ExpectedDBs, nil, nil
}

func (m *MockGodoDatabaseSvc) ListDBs(context.Context, string, *godo.ListOptions) ([]godo.DatabaseDB, *godo.Response, error) {
	return ExpectedDatabaseDBs, nil, nil

}

func (m *MockGodoDatabaseSvc) Migrate(context.Context, string, *godo.DatabaseMigrateRequest) (*godo.Response, error) {
	return nil, errors.New(TestError)

}

func (m *MockGodoDatabaseSvc) Resize(context.Context, string, *godo.DatabaseResizeRequest) (*godo.Response, error) {
	return nil, errors.New(TestError)
}

func (m *MockGodoDatabaseSvc) UpdateMaintenance(context.Context, string, *godo.DatabaseUpdateMaintenanceRequest) (*godo.Response, error) {
	return nil, errors.New(TestError)
}
