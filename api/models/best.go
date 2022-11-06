package models

import (
	"math"
	"sync"
)

type Best struct {
	M               sync.RWMutex
	beta            int
	beta_percentage int
}

// Get gets the beta value using the RWMutex
func (b *Best) Get() (beta int) {
	b.M.RLock()
	beta = b.beta
	b.M.RUnlock()

	return
}

// GetPercentage gets the beta percentage value using the RWMutex
func (b *Best) GetPercentage() (beta_percentage int) {
	b.M.RLock()
	beta_percentage = b.beta_percentage
	b.M.RUnlock()

	return
}

// GetAll gets the beta and his percentage value using the RWMutex
func (b *Best) GetAll() (beta, beta_percentage int) {
	b.M.RLock()
	beta = b.beta
	beta_percentage = b.beta_percentage
	b.M.RUnlock()

	return
}

// Set sets the beta value and pre-compute his percentage using the RWMutex
func (b *Best) Set(beta int) {
	b.M.Lock()
	b.beta = beta
	b.beta_percentage = beta * 100
	b.M.Unlock()

	return
}

// SetAndGetPercentage sets the beta value, pre-compute his percentage and return this one using the RWMutex
func (b *Best) SetAndGetPercentage(beta int) (beta_percentage int) {
	beta_percentage = beta * 100

	b.M.Lock()
	b.beta = beta
	b.beta_percentage = beta_percentage
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

// MaxAndGetPercentage replaces the alpha value if is lower than the beta value, and return the max and his percentage, using the RWMutex
func (b *Best) MaxAndGetAll(beta int) (max, max_percentage int, first bool) {
	max, max_percentage = b.GetAll()

	if max == math.MinInt {
		first = true
	} else if beta <= max {

		return
	}

	max = beta
	max_percentage = b.SetAndGetPercentage(max)

	return
}
