package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/gen/knowledge"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/internal/repositories/mocks"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
)

// Simple test to verify mock setup
func TestMockSetup(t *testing.T) {
	// Create mock
	profilesRepo := &mocks.MockProfilesServiceRepository{}

	// Setup mock expectation
	expectedProfile := createTestCompleteProfile("student")
	profilesRepo.On("GetCompleteProfile", mock.Anything, &profiles.GetCompleteProfilePayload{
		SessionToken: "valid_token",
	}).Return(expectedProfile, nil)

	// Call the mocked method
	result, err := profilesRepo.GetCompleteProfile(context.Background(), &profiles.GetCompleteProfilePayload{
		SessionToken: "valid_token",
	})

	// Verify
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProfile.Role, result.Role)

	// Verify all expectations were met
	profilesRepo.AssertExpectations(t)
}

// Test simple CreateTest to verify it works
func TestCreateTestSimple(t *testing.T) {
	// Setup mocks
	testRepo := &mocks.MockTestRepository{}
	profilesRepo := &mocks.MockProfilesServiceRepository{}

	// Setup expectations
	expectedProfile := createTestCompleteProfile("teacher")
	profilesRepo.On("GetCompleteProfile", mock.Anything, mock.AnythingOfType("*profiles.GetCompleteProfilePayload")).Return(expectedProfile, nil)
	testRepo.On("CreateTest", mock.Anything, mock.AnythingOfType("knowledgedb.CreateTestParams")).Return(nil)

	// Create service
	service := setupTestService(testRepo, nil, nil, nil, profilesRepo)

	// Call method
	result, err := service.CreateTest(context.Background(), &knowledge.CreateTestPayload{
		SessionToken: "valid_token",
		Title:        "Test Quiz",
	})

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Equal(t, "Test created successfully", result.Message)

	// Verify all expectations were met
	testRepo.AssertExpectations(t)
	profilesRepo.AssertExpectations(t)
}
