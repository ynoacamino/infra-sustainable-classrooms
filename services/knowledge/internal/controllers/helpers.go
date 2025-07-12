package controllers

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Helper to convert int64 timestamp to pgtype.Timestamptz
func (s *knowledgesvrc) IntToPgtimestamptz(i int64) pgtype.Timestamptz {
	if i == 0 {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{
		Time:  time.UnixMilli(i),
		Valid: true,
	}
}

// Helper to convert pgtype.Timestamptz to int64 timestamp
func (s *knowledgesvrc) PgtimestamptzToInt64(ts pgtype.Timestamptz) int64 {
	if !ts.Valid {
		return 0
	}
	return ts.Time.UnixMilli()
}

// Helper to convert int32 to pgtype.Int4
func (s *knowledgesvrc) Int32ToPgInt4(i int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: i,
		Valid: true,
	}
}

// Helper to convert pgtype.Int4 to int32
func (s *knowledgesvrc) PgInt4ToInt32(i pgtype.Int4) int32 {
	if !i.Valid {
		return 0
	}
	return i.Int32
}

// Helper to convert bool to pgtype.Bool
func (s *knowledgesvrc) BoolToPgBool(b bool) pgtype.Bool {
	return pgtype.Bool{
		Bool:  b,
		Valid: true,
	}
}

// Helper to convert pgtype.Bool to bool
func (s *knowledgesvrc) PgBoolToBool(b pgtype.Bool) bool {
	if !b.Valid {
		return false
	}
	return b.Bool
}

// Helper to convert float64 to pgtype.Numeric
func (s *knowledgesvrc) Float64ToPgNumeric(f float64) pgtype.Numeric {
	var num pgtype.Numeric
	err := num.Scan(f)
	if err != nil {
		// If direct scan fails, try string conversion
		err = num.Scan(fmt.Sprintf("%f", f))
		if err != nil {
			return pgtype.Numeric{Valid: false}
		}
	}
	return num
}

// Helper to convert pgtype.Numeric to float64
func (s *knowledgesvrc) PgNumericToFloat64(num pgtype.Numeric) float64 {
	if !num.Valid {
		return 0.0
	}
	f, err := num.Float64Value()
	if err != nil {
		return 0.0
	}
	return f.Float64
}
