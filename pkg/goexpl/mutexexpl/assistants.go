package mutexexpl

import "sync"

// MutexContainer
// a MutexContainer is used to hold a map of counters and a mutex, which allows multiple goroutines to update counters
// concurrently
// mutex cannot be copied, so we need to pass a pointer when passing MutexContainer around
type MutexContainer struct {
	mu       sync.Mutex
	counters map[string]int
}

func (mc *MutexContainer) inc(name string) {
	// lock before operating on the map counters
	mc.mu.Lock()
	defer mc.mu.Unlock() // unlock in defer to ensure it is always unlocked
	mc.counters[name]++
}
