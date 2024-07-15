package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func PgTimestampGetter(time *time.Time) pgtype.Timestamp {
	if time == nil {
		return pgtype.Timestamp{}
	}

	return pgtype.Timestamp{
		Time: *time,
	}
}
