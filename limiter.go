package RateLimiter

import (
	"math"
	"sync"
	"time"
)

type Limiter struct {
	window time.Duration
	limit  int
	tally  []time.Time
	mux    sync.RWMutex
	//todo test mutex
}

func (l *Limiter) RemainingTallyAsGroup2() [][]time.Time {
	return [][]time.Time{l.tally}
}

func NewLimiter(Window time.Duration, Units int) ILimiter {
	return &Limiter{
		window: Window,
		limit:  Units,
	}
}

func (l *Limiter) AddTally() bool {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.tally = append(l.tally, time.Now())
	return len(l.tally) < l.limit
}

func (l *Limiter) TryAddTally() bool {
	if l.RemainingTally() > 0 {
		l.AddTally()
		return true
	}
	return false
}

func (l *Limiter) HasRemainingTally() bool {
	l.CleanUp()
	return l.limit-l.len() >= 1
}

func (l *Limiter) RemainingTally() int {
	l.CleanUp()
	return l.limit - l.len()
}

func (l *Limiter) len() int {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return len(l.tally)
}

func (l *Limiter) CleanUp() {
	l.mux.Lock()
	defer l.mux.Unlock()
	for i, tv := range l.tally {
		ts := time.Now().Sub(tv)
		if ts > l.window {
			if len(l.tally) == 1 {
				l.tally = []time.Time{}
				continue
			}
			low := int(math.Max(float64(i), 0))
			high := int(math.Min(float64(i+1), float64(len(l.tally))))
			l.tally = append(l.tally[:low], l.tally[high:]...)
		}
	}
}

func (l *Limiter) Window() time.Duration {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return l.window
}

func (l *Limiter) GetStorableData() interface{} {
	var res []string
	for _, t := range l.tally {
		res = append(res, t.Format(time.RFC3339))
	}
	return res
}

func (l *Limiter) SetStorableData(data interface{}) bool {
	var res []time.Time
	rsc := true
	itd, ok := data.([]interface{})
	if !ok {
		return false
	}
	for _, td := range itd {
		ts, ok := td.(string)
		if !ok {
			rsc = false
			continue
		}
		tp, err := time.Parse(time.RFC3339, ts)
		if err != nil {
			rsc = false
			continue
		}
		res = append(res, tp)
	}

	l.tally = res
	return rsc
}

func (l *Limiter) GetDebug() interface{} {
	return DebugInfo{
		Window: l.window,
		Limit:  l.limit,
		Tally:  l.tally,
	}
}
