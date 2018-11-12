package pipelines

import (
	"time"
)

// Project represents the spec of a project pipeline.
type Project struct {
	PollInterval time.Duration
}
