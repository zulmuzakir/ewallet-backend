package helper

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func TimeToTimestamptz(t time.Time) pgtype.Timestamptz {
	if t.IsZero() {
		return pgtype.Timestamptz{Valid: false}
	}

	return pgtype.Timestamptz{Valid: true, Time: t}
}
