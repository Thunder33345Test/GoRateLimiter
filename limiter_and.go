package RateLimiter

import (
	"math"
)

type LimiterGroupAnd struct {
	limiters []ILimiter
}

func NewLimiterGroupAnd(limiters ...ILimiter) ILimiter {
	return &LimiterGroupAnd{limiters: limiters}
}

func (l LimiterGroupAnd) AddTally() bool {

	for _, ls := range l.limiters {
		ls.AddTally()
	}
	return l.HasRemainingTally()
}

func (l LimiterGroupAnd) TryAddTally() bool {
	if !l.HasRemainingTally() {
		return false
	}
	for _, ls := range l.limiters {
		ls.AddTally()
	}
	return true
}

func (l LimiterGroupAnd) HasRemainingTally() bool {
	for _, ls := range l.limiters {
		if !ls.HasRemainingTally() {
			return false
		}
	}
	return true
}

func (l LimiterGroupAnd) RemainingTally() int {
	least := math.MaxInt64
	for _, ls := range l.limiters {
		rt := ls.RemainingTally()
		if rt < least {
			least = rt
		}
	}
	return least
}

func (l LimiterGroupAnd) CleanUp() {
	for _, ls := range l.limiters {
		ls.CleanUp()
	}
}

func (l LimiterGroupAnd) GetGroups() []ILimiter {
	return l.limiters
}

func (l LimiterGroupAnd) RemainingTallyAsGroup() []int {
	var s []int
	for _, ls := range l.limiters {
		s = append(s, ls.RemainingTally())
	}
	return s
}

func (l *LimiterGroupAnd) GetStorableData() interface{} {
	return commonGetStorable(l)
}

func (l *LimiterGroupAnd) SetStorableData(i interface{}) bool {
	return commonSetStorable(l, i)
}

func (l *LimiterGroupAnd) GetDebug() interface{} {
	return commonGetDebug(l,"AND")
}
