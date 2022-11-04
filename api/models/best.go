package models

import (
	"math"
	"sync"
)

type Best struct {
	M    sync.RWMutex
	beta int
}

// Get gets the beta value using the RWMutex
func (b *Best) Get() (beta int) {
	b.M.RLock()
	beta = b.beta
	b.M.RUnlock()

	return
}

// Set sets the beta value using the RWMutex
func (b *Best) Set(beta int) {
	b.M.Lock()
	b.beta = beta
	b.M.Unlock()

	return
}

// Max replaces the alpha value if is lower than the beta value, and return the max, using the RWMutex
func (b *Best) Max(beta int) (max int, first bool) {
	max = b.Get()

	if max == math.MinInt {
		first = true
	} else if beta <= max {
		return
	}

	b.Set(beta)
	max = beta

	return
}
