package ch1_5

import "sync"

// use sync.Once realize singleton
var (
	instance_ *singleton
	once      sync.Once
)

func Instance_() *singleton {
	once.Do(func() {
		instance_ = &singleton{}
	})
	return instance
}
