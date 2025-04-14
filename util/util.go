package util

import (
	"context"
	"time"

	"github.com/0xPolygon/panoptichain/log"
)

// BlockFor pauses execution until the current time is rounded to the nearest
// multiple of the specified duration, or until the context is canceled.
func BlockFor(ctx context.Context, duration time.Duration) {
	now := time.Now()
	rounded := now.Add(duration / 2).Round(duration)

	log.Trace().Time("now", now).Time("until", rounded).Msg("Blocking")

	timer := time.NewTimer(time.Until(rounded))
	defer timer.Stop()

	select {
	case <-timer.C:
		return
	case <-ctx.Done():
		return
	}
}
