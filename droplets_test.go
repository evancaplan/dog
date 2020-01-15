package dog

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/digitalocean/godo"
)

var TestDroplet = godo.Droplet{
	ID:     1,
	Name:   "test.example.com",
	Memory: 1024,
	Vcpus:  1,
	Disk:   25,
	Locked: true,
	Status: "new",
	Kernel: &godo.Kernel{
		ID:      1,
		Name:    "Ubuntu BWSS Test Droplet Name",
		Version: "3.13.0-37-generic",
	},
	Created:     "2020-1-1T16:36:31Z",
	Features:    []string{"virtio"},
	BackupIDs:   []int{1, 2, 3, 4},
	SnapshotIDs: []int{1, 2, 4},
	Image:       &godo.Image{},
	VolumeIDs:   []string{"2", "3", "4"},
	Size:        &godo.Size{},
	SizeSlug:    "s-1vcpu-1gb",
	Networks:    &godo.Networks{},
	Region:      &godo.Region{},
	Tags:        []string{"tag"},
}

var TestDroplets = []godo.Droplet{TestDroplet}

var TestFindAllDropletsRequest = FindAllDropletsRequest{
	Page:    1,
	PerPage: 5,
}

var TestFindDropletByIDRequest = FindDropletByIDRequest{
	ID: 1234,
}

var TestFindDropletsByTagRequest = FindDropletsByTagRequest{
	Tag:     "Test Tag",
	Page:    1,
	PerPage: 5,
}

var TestCreateDropletRequest = CreateDropletRequest{
	Name:              "Droplet Name",
	Region:            NYC2,
	DropletSize:       S3Cpu1GbRAM,
	Image:             "Test Image",
	SSHKeys:           []int{1, 2, 3},
	Backups:           true,
	IPv6:              false,
	Configuration:     "Test Config",
	PrivateNetworking: false,
	Volumes:           []string{"test", "volumes"},
	Tags:              []string{"dog"},
	VPCUUID:           "ASD-342",
}

var TestDeleteDropletRequest = DeleteDropletRequest{
	ID: 1,
}

func TestGetAllDroplets(t *testing.T) {

	dbClient := NewDC(TestPAT)
	dbClient.client = &MockGodoDropletSvc{}

	expected := TestDroplets
	returned, _ := dbClient.GetAllDroplets(TestFindAllDropletsRequest)
	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("expected %+v\n , returned, %+v\n ", expected, returned)
	}

}

func TestGetDropletById(t *testing.T) {

	dbClient := NewDC(TestPAT)
	dbClient.client = &MockGodoDropletSvc{}

	expected := &TestDroplet
	returned, _ := dbClient.GetDropletById(TestFindDropletByIDRequest)
	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("expected %+v\n , returned, %+v\n ", expected, returned)
	}

}

func TestGetDropletsByTag(t *testing.T) {

	dbClient := NewDC(TestPAT)
	dbClient.client = &MockGodoDropletSvc{}

	expected := &TestDroplets
	returned, _ := dbClient.GetDropletsByTag(TestFindDropletsByTagRequest)
	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("expected %+v\n , returned, %+v\n ", expected, returned)
	}

}

func TestCreateDroplet(t *testing.T) {

	dbClient := NewDC(TestPAT)
	dbClient.client = &MockGodoDropletSvc{}

	expected := &TestDroplet
	returned, _ := dbClient.CreateDroplet(TestCreateDropletRequest)
	if !reflect.DeepEqual(expected, returned) {
		t.Errorf("expected %+v\n , returned, %+v\n ", expected, returned)
	}

}

func TestDeleteDroplet(t *testing.T) {

	t.Run("Error is thrown", func(t *testing.T) {
		dbClient := NewDC(TestPAT)
		dbClient.client = &MockGodoDropletSvc{}

		expectedError := "Unable to delete droplet with ID: " + strconv.Itoa(TestDeleteDropletRequest.ID)
		returnedError := dbClient.DeleteDroplet(TestDeleteDropletRequest)
		if expectedError != returnedError.Error() {
			t.Errorf("expected error: %s returned error: %s", expectedError, returnedError)
		}
	})

}

type MockGodoDropletSvc struct{}

func (m *MockGodoDropletSvc) List(context.Context, *godo.ListOptions) ([]godo.Droplet, *godo.Response, error) {
	return TestDroplets, nil, nil
}

func (m *MockGodoDropletSvc) ListByTag(context.Context, string, *godo.ListOptions) ([]godo.Droplet, *godo.Response, error) {
	return TestDroplets, nil, nil
}

func (m *MockGodoDropletSvc) Get(context.Context, int) (*godo.Droplet, *godo.Response, error) {
	return &TestDroplet, nil, nil
}

func (m *MockGodoDropletSvc) Create(context.Context, *godo.DropletCreateRequest) (*godo.Droplet, *godo.Response, error) {
	return &TestDroplet, nil, nil
}

func (m *MockGodoDropletSvc) Delete(context.Context, int) (*godo.Response, error) {
	return nil, errors.New("Test Godo Error")
}
