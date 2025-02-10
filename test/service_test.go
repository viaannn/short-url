package unit_test

import (
	"short_url/internal/app"
	unit "short_url/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepo = new(unit.RepositoryMock)
var service = app.InitService(mockRepo)

// TestCreateShortLink_ReturnSuccess tests the CreateShortLink method in the service
func TestCreateShortLink_ReturnSuccess(t *testing.T) {
	mockRepo.Mock.On("CreateMasterUrl", mock.Anything, "http://example.com").Return("key", "http://example.com", nil)

	// Call the service method
	request := app.CreateShortLinkRequest{
		TargetUrl: "http://example.com",
	}

	// Create master URL
	response, err := service.CreateMasterUrl(request)
	if err != nil {
		t.Errorf("Error was not expected, got: %v", err)
	}

	// Assert the response
	assert.Equal(t, "http://localhost:8080/key", response.ShortLink)
	assert.Equal(t, "http://example.com", response.TargetUrl)
}

// TestRedirectShortLink_ReturnSuccess tests the GetTargetUrl method in the service
func TestRedirectShortLink_ReturnSuccess(t *testing.T) {
	mockRepo.Mock.On("FindByKey", mock.Anything).Return("key", "http://example.com", nil)

	// Call the service method
	targetUrl, err := service.GetTargetUrl("key")
	if err != nil {
		t.Errorf("Error was not expected, got: %v", err)
	}

	// Assert the response
	assert.Equal(t, "http://example.com", targetUrl)
}
