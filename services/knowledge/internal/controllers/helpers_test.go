package controllers

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestIntToPgtimestamptz(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name     string
		input    int64
		expected pgtype.Timestamptz
	}{
		{
			name:     "zero timestamp",
			input:    0,
			expected: pgtype.Timestamptz{Time: time.Time{}, Valid: false},
		},
		{
			name:  "valid timestamp",
			input: 1609459200000, // 2021-01-01 00:00:00 UTC in milliseconds
			expected: pgtype.Timestamptz{
				Time:  time.UnixMilli(1609459200000),
				Valid: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.IntToPgtimestamptz(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPgtimestamptzToInt64(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name     string
		input    pgtype.Timestamptz
		expected int64
	}{
		{
			name:     "invalid timestamp",
			input:    pgtype.Timestamptz{Time: time.Time{}, Valid: false},
			expected: 0,
		},
		{
			name: "valid timestamp",
			input: pgtype.Timestamptz{
				Time:  time.UnixMilli(1609459200000),
				Valid: true,
			},
			expected: 1609459200000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.PgtimestamptzToInt64(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInt32ToPgInt4(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name     string
		input    int32
		expected pgtype.Int4
	}{
		{
			name:     "zero value",
			input:    0,
			expected: pgtype.Int4{Int32: 0, Valid: true},
		},
		{
			name:     "positive value",
			input:    42,
			expected: pgtype.Int4{Int32: 42, Valid: true},
		},
		{
			name:     "negative value",
			input:    -10,
			expected: pgtype.Int4{Int32: -10, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.Int32ToPgInt4(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPgInt4ToInt32(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name     string
		input    pgtype.Int4
		expected int32
	}{
		{
			name:     "invalid value",
			input:    pgtype.Int4{Int32: 0, Valid: false},
			expected: 0,
		},
		{
			name:     "valid value",
			input:    pgtype.Int4{Int32: 42, Valid: true},
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.PgInt4ToInt32(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBoolToPgBool(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name     string
		input    bool
		expected pgtype.Bool
	}{
		{
			name:     "true value",
			input:    true,
			expected: pgtype.Bool{Bool: true, Valid: true},
		},
		{
			name:     "false value",
			input:    false,
			expected: pgtype.Bool{Bool: false, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.BoolToPgBool(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPgBoolToBool(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name     string
		input    pgtype.Bool
		expected bool
	}{
		{
			name:     "invalid value",
			input:    pgtype.Bool{Bool: false, Valid: false},
			expected: false,
		},
		{
			name:     "valid true value",
			input:    pgtype.Bool{Bool: true, Valid: true},
			expected: true,
		},
		{
			name:     "valid false value",
			input:    pgtype.Bool{Bool: false, Valid: true},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.PgBoolToBool(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFloat64ToPgNumeric(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name  string
		input float64
	}{
		{
			name:  "zero value",
			input: 0.0,
		},
		{
			name:  "positive value",
			input: 42.5,
		},
		{
			name:  "negative value",
			input: -10.25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.Float64ToPgNumeric(tt.input)
			assert.True(t, result.Valid)

			// Convert back to verify
			converted := service.PgNumericToFloat64(result)
			assert.Equal(t, tt.input, converted)
		})
	}
}

func TestPgNumericToFloat64(t *testing.T) {
	service := &knowledgesvrc{}

	tests := []struct {
		name     string
		input    pgtype.Numeric
		expected float64
	}{
		{
			name:     "invalid value",
			input:    pgtype.Numeric{Valid: false},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.PgNumericToFloat64(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
