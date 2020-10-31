package RateLimiter

import "time"

type ILimiter interface {
	AddTally() bool
	TryAddTally() bool
	HasRemainingTally() bool
	RemainingTally() int
	CleanUp()
}

type ILimiterGroup interface {
	GetGroups() []ILimiter
	RemainingTallyAsGroup() []int
}

type Storable interface {
	GetStorableData() interface{}
	SetStorableData(interface{}) bool
}

type Debuggable interface {
	GetDebug() interface{}
}

type DebugInfo struct {
	Window time.Duration
	Limit  int
	Tally  []time.Time
}
