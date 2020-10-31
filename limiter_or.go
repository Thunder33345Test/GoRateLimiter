package RateLimiter

type LimiterGroupOr struct {
	limiters []ILimiter
}

func NewLimiterGroupOr(limiters ...ILimiter) ILimiter {
	return &LimiterGroupOr{limiters: limiters}
}

func (l LimiterGroupOr) AddTally() bool {
	for _, ls := range l.limiters {
		if ls.TryAddTally() {
			return l.HasRemainingTally()
		}
	}
	if v := l.limiters[0]; v != nil {
		v.AddTally()
	}
	return l.HasRemainingTally()
}

func (l LimiterGroupOr) TryAddTally() bool {
	if !l.HasRemainingTally() {
		return false
	}
	for _, ls := range l.limiters {
		if ls.TryAddTally() {
			return true
		}
	}
	return false
}

func (l LimiterGroupOr) HasRemainingTally() bool {
	for _, ls := range l.limiters {
		if ls.HasRemainingTally() {
			return true
		}
	}
	return false
}

func (l LimiterGroupOr) RemainingTally() int {
	rem := 0
	for _, ls := range l.limiters {
		rem += ls.RemainingTally()
	}
	return rem
}

func (l LimiterGroupOr) CleanUp() {
	for _, ls := range l.limiters {
		ls.CleanUp()
	}
}

func (l LimiterGroupOr) GetGroups() []ILimiter {
	return l.limiters
}

func (l LimiterGroupOr) RemainingTallyAsGroup() []int {
	var s []int
	for _, ls := range l.limiters {
		s = append(s, ls.RemainingTally())
	}
	return s
}

func (l *LimiterGroupOr) GetStorableData() interface{} {
	return commonGetStorable(l)
}

func (l *LimiterGroupOr) SetStorableData(i interface{}) bool {
	return commonSetStorable(l, i)
}

func (l *LimiterGroupOr) GetDebug() interface{} {
	return commonGetDebug(l, "OR")
}
