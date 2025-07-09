// Package config provides helper functions for parsing environment variables
// with fallback defaults and type conversions.
package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

// getEnvOrDefault returns the environment variable value or a default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("Environment variable %s not found, using default value: %s\n", key, defaultValue)
	return defaultValue
}

// parseBoolOrDefault parses a boolean environment variable or returns a default value
func parseBoolOrDefault(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
		log.Printf("Invalid boolean value for %s: %s, using default: %t\n", key, value, defaultValue)
	}
	return defaultValue
}

// parseIntOrDefault parses an integer environment variable or returns a default value
func parseIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
		log.Printf("Invalid integer value for %s: %s, using default: %d\n", key, value, defaultValue)
	}
	return defaultValue
}

// parseDurationOrDefault parses a duration environment variable or returns a default value
func parseDurationOrDefault(key string, defaultValue string) time.Duration {
	value := getEnvOrDefault(key, defaultValue)
	if duration, err := time.ParseDuration(value); err == nil {
		return duration
	}
	// Fallback to default if parsing fails
	if duration, err := time.ParseDuration(defaultValue); err == nil {
		log.Printf("Invalid duration value for %s: %s, using default: %s\n", key, value, defaultValue)
		return duration
	}
	// Last resort fallback
	log.Printf("Critical error: cannot parse default duration %s, using 1 hour\n", defaultValue)
	return time.Hour
}
