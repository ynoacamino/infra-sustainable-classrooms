package controllers

import (
	"context"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	auth "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/auth"
	authdb "github.com/ynoacamino/infra-sustainable-classrooms/services/auth/gen/database"
	"github.com/ynoacamino/infra-sustainable-classrooms/services/auth/internal/repositories/mocks"
	"golang.org/x/crypto/bcrypt"
)

func TestGetSessionExpirationPG(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	// Act
	result := service.getSessionExpirationPG()

	// Assert
	assert.True(t, result.Valid)
	assert.True(t, result.Time.After(time.Now()))
	assert.True(t, result.Time.Before(time.Now().Add(25*time.Hour))) // Should be ~24h from now
}

func TestGenerateSessionToken(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	// Act
	token, err := service.generateSessionToken()

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, 64, len(token)) // 32 bytes = 64 hex chars

	// Test that tokens are unique
	token2, err2 := service.generateSessionToken()
	assert.NoError(t, err2)
	assert.NotEqual(t, token, token2)
}

func TestGenerateBackupCodes(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	// Act
	codes, err := service.generateBackupCodes(10)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, codes, 10)

	// Check that all codes are unique
	codeSet := make(map[string]bool)
	for _, code := range codes {
		assert.False(t, codeSet[code], "Duplicate code found: %s", code)
		codeSet[code] = true
		assert.Equal(t, 8, len(code)) // Default length is 8

		// Check that code contains only valid characters
		matched, _ := regexp.MatchString("^[A-Z0-9]+$", code)
		assert.True(t, matched, "Code contains invalid characters: %s", code)
	}
}

func TestGenerateBackupCodes_ZeroCount(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	// Act
	codes, err := service.generateBackupCodes(0)

	// Assert
	assert.NoError(t, err)
	assert.Empty(t, codes)
}

func TestGenerateRandomCode(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	testCases := []struct {
		name   string
		length int
	}{
		{"Length 1", 1},
		{"Length 8", 8},
		{"Length 16", 16},
		{"Length 32", 32},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			code, err := service.generateRandomCode(tc.length)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.length, len(code))

			// Check that code contains only valid characters
			matched, _ := regexp.MatchString("^[A-Z0-9]+$", code)
			assert.True(t, matched, "Code contains invalid characters: %s", code)
		})
	}
}

func TestHashBackupCode(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)
	testCode := "TESTCODE123"

	// Act
	hash := service.hashBackupCode(testCode)

	// Assert
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, testCode, hash)
	assert.True(t, strings.HasPrefix(hash, "$2a$") || strings.HasPrefix(hash, "$2b$"), "Hash should be bcrypt format")

	// Test that same code produces different hashes (due to salt)
	hash2 := service.hashBackupCode(testCode)
	assert.NotEqual(t, hash, hash2)
}

func TestVerifyBackupCode(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)
	testCode := "TESTCODE123"

	// Create a valid hash
	hash, err := bcrypt.GenerateFromPassword([]byte(testCode), bcrypt.DefaultCost)
	assert.NoError(t, err)

	// Test valid code
	t.Run("Valid Code", func(t *testing.T) {
		result := service.verifyBackupCode(testCode, string(hash))
		assert.True(t, result)
	})

	// Test invalid code
	t.Run("Invalid Code", func(t *testing.T) {
		result := service.verifyBackupCode("WRONGCODE", string(hash))
		assert.False(t, result)
	})

	// Test invalid hash
	t.Run("Invalid Hash", func(t *testing.T) {
		result := service.verifyBackupCode(testCode, "invalid-hash")
		assert.False(t, result)
	})
}

func TestGetDeviceInfoText(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	userAgent := "Mozilla/5.0"
	deviceID := "device-123"
	platform := "web"

	deviceInfo := &auth.DeviceInfo{
		UserAgent: &userAgent,
		DeviceID:  &deviceID,
		Platform:  &platform,
	}

	testCases := []struct {
		name     string
		field    string
		expected string
	}{
		{"User Agent", "user_agent", userAgent},
		{"Device ID", "device_id", deviceID},
		{"Platform", "platform", platform},
		{"Invalid Field", "invalid", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := service.getDeviceInfoText(deviceInfo, tc.field)

			// Assert
			if tc.expected == "" {
				assert.False(t, result.Valid)
				assert.Equal(t, "", result.String)
			} else {
				assert.True(t, result.Valid)
				assert.Equal(t, tc.expected, result.String)
			}
		})
	}
}

func TestGetDeviceInfoText_NilDeviceInfo(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	// Act
	result := service.getDeviceInfoText(nil, "user_agent")

	// Assert
	assert.False(t, result.Valid)
	assert.Equal(t, "", result.String)
}

func TestGetDeviceInfoText_NilFields(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	// DeviceInfo with nil fields
	deviceInfo := &auth.DeviceInfo{
		UserAgent: nil,
		DeviceID:  nil,
		Platform:  nil,
	}

	// Act
	result := service.getDeviceInfoText(deviceInfo, "user_agent")

	// Assert
	assert.False(t, result.Valid)
	assert.Equal(t, "", result.String)
}

func TestGetDeviceInfoIP(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	t.Run("Valid IPv4", func(t *testing.T) {
		ipStr := "192.168.1.1"
		deviceInfo := &auth.DeviceInfo{
			IPAddress: &ipStr,
		}

		// Act
		result := service.getDeviceInfoIP(deviceInfo)

		// Assert
		assert.NotNil(t, result)
		assert.Equal(t, "192.168.1.1", result.String())
	})

	t.Run("Valid IPv6", func(t *testing.T) {
		ipStr := "2001:db8::1"
		deviceInfo := &auth.DeviceInfo{
			IPAddress: &ipStr,
		}

		// Act
		result := service.getDeviceInfoIP(deviceInfo)

		// Assert
		assert.NotNil(t, result)
		assert.Equal(t, "2001:db8::1", result.String())
	})

	t.Run("Invalid IP", func(t *testing.T) {
		ipStr := "invalid-ip"
		deviceInfo := &auth.DeviceInfo{
			IPAddress: &ipStr,
		}

		// Act
		result := service.getDeviceInfoIP(deviceInfo)

		// Assert
		assert.Nil(t, result)
	})

	t.Run("Nil DeviceInfo", func(t *testing.T) {
		// Act
		result := service.getDeviceInfoIP(nil)

		// Assert
		assert.Nil(t, result)
	})

	t.Run("Nil IP Address", func(t *testing.T) {
		deviceInfo := &auth.DeviceInfo{
			IPAddress: nil,
		}

		// Act
		result := service.getDeviceInfoIP(deviceInfo)

		// Assert
		assert.Nil(t, result)
	})
}

func TestValidateSession(t *testing.T) {
	t.Run("Valid Session", func(t *testing.T) {
		// Arrange
		mockSessionRepo := &mocks.MockSessionRepository{}
		service := setupTestService(nil, mockSessionRepo, nil, nil, nil)

		validSession := authdb.GetSessionByTokenRow{
			ID:           1,
			UserID:       1,
			SessionToken: "valid-token",
			IsActive:     true,
			ExpiresAt: pgtype.Timestamptz{
				Time:  time.Now().Add(1 * time.Hour),
				Valid: true,
			},
		}

		mockSessionRepo.On("GetSessionByToken", mock.Anything, "valid-token").Return(validSession, nil)

		// Act
		result, err := service.validateSession(context.Background(), "valid-token")

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int64(1), result.UserID)
		assert.True(t, result.IsActive)

		mockSessionRepo.AssertExpectations(t)
	})

	t.Run("Empty Token", func(t *testing.T) {
		// Arrange
		service := setupTestService(nil, nil, nil, nil, nil)

		// Act
		result, err := service.validateSession(context.Background(), "")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "session token is required")
	})

	t.Run("Session Not Found", func(t *testing.T) {
		// Arrange
		mockSessionRepo := &mocks.MockSessionRepository{}
		service := setupTestService(nil, mockSessionRepo, nil, nil, nil)

		mockSessionRepo.On("GetSessionByToken", mock.Anything, "invalid-token").Return(authdb.GetSessionByTokenRow{}, assert.AnError)

		// Act
		result, err := service.validateSession(context.Background(), "invalid-token")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "session not found")

		mockSessionRepo.AssertExpectations(t)
	})

	t.Run("Inactive Session", func(t *testing.T) {
		// Arrange
		mockSessionRepo := &mocks.MockSessionRepository{}
		service := setupTestService(nil, mockSessionRepo, nil, nil, nil)

		inactiveSession := authdb.GetSessionByTokenRow{
			ID:           1,
			UserID:       1,
			SessionToken: "inactive-token",
			IsActive:     false,
			ExpiresAt: pgtype.Timestamptz{
				Time:  time.Now().Add(1 * time.Hour),
				Valid: true,
			},
		}

		mockSessionRepo.On("GetSessionByToken", mock.Anything, "inactive-token").Return(inactiveSession, nil)

		// Act
		result, err := service.validateSession(context.Background(), "inactive-token")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "session is inactive")

		mockSessionRepo.AssertExpectations(t)
	})

	t.Run("Expired Session", func(t *testing.T) {
		// Arrange
		mockSessionRepo := &mocks.MockSessionRepository{}
		service := setupTestService(nil, mockSessionRepo, nil, nil, nil)

		expiredSession := authdb.GetSessionByTokenRow{
			ID:           1,
			UserID:       1,
			SessionToken: "expired-token",
			IsActive:     true,
			ExpiresAt: pgtype.Timestamptz{
				Time:  time.Now().Add(-1 * time.Hour), // Expired
				Valid: true,
			},
		}

		mockSessionRepo.On("GetSessionByToken", mock.Anything, "expired-token").Return(expiredSession, nil)

		// Act
		result, err := service.validateSession(context.Background(), "expired-token")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "session already expired")

		mockSessionRepo.AssertExpectations(t)
	})
}

func TestTruncateToken(t *testing.T) {
	// Arrange
	service := setupTestService(nil, nil, nil, nil, nil)

	testCases := []struct {
		name     string
		token    string
		expected string
	}{
		{"Short Token", "abc", "a..."},
		{"Medium Token", "abcdefgh", "abcd..."},
		{"Long Token", "abcdefghijklmnopqrstuvwxyz", "abcdefghij..."},
		{"Exactly 10 chars", "abcdefghij", "abcde..."},
		{"Empty Token", "", "..."},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := service.truncateToken(tc.token)

			// Assert
			assert.Equal(t, tc.expected, result)
		})
	}
}
