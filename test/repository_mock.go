package unit

import (
	"short_url/internal/app"

	"github.com/stretchr/testify/mock"
)

// RepositoryMock is a mock of Repository interface or component
type RepositoryMock struct {
	mock.Mock
}

// CreateShortLink is a mock of CreateShortLink method in Repository interface
// If args are valid, it returns a MasterUrl mock object and nil error
// If args not valid or nil, it returns nil and an error
func (repo *RepositoryMock) CreateMasterUrl(key string, targetUrl string) (*app.MasterUrl, error) {
	// Call the mock method
	args := repo.Mock.Called(key, targetUrl)
	if args.Get(0) == nil || args.Get(1) == nil {
		return nil, args.Error(1)
	}

	// Create a mock MasterUrl object
	mockMasterUrl := &app.MasterUrl{
		Key:       args.Get(0).(string),
		TargetUrl: args.Get(1).(string),
	}

	return mockMasterUrl, nil
}

// FindByKey is a mock of FindByKey method in Repository interface
// If args key are valid, it returns a MasterUrl mock object and nil error
func (repo *RepositoryMock) FindByKey(key string) (*app.MasterUrl, error) {
	args := repo.Mock.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	mockMasterUrl := &app.MasterUrl{
		Key:       args.Get(0).(string),
		TargetUrl: "http://example.com",
	}

	return mockMasterUrl, nil
}
