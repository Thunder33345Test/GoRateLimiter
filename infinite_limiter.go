package RateLimiter

import (
	"math"
)

type InfiniteLimiter struct {
}

func NewInfiniteLimiter() *InfiniteLimiter {
	return &InfiniteLimiter{}
}

func (i InfiniteLimiter) AddTally() bool {
	return true
}

func (i InfiniteLimiter) TryAddTally() bool {
	return true
}

func (i InfiniteLimiter) RemainingTally() int {
	return math.MaxUint16
}

func (i InfiniteLimiter) HasRemainingTally() bool {
	return true
}

func (i InfiniteLimiter) CleanUp() {

}