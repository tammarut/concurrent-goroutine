package main

import (
	"github.com/tammarut/concurrent-goroutine/solution7"
)

func main() {
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 1: Basic Goroutines
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// solution1.SimpleGoroutine()

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 2: WaitGroups
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// solution2.WaitGroups()

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 3: Channels
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// solution3.UseChannels()

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 4: Worker Pool
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// solution4.WorkerPool()

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 5: Limiting Goroutines with channels (use channels to create a semaphore-like mechanism)
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// solution5.LimitGoroutinesWithChannel()

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 6: Limiting Goroutines with library sync/semaphore
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// solution6.LimitGoroutinesWithSemaphore()

	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	// Solution 7: Library errgroup
	// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
	solution7.LimitGoroutinesWithLibrarySemaphore()
}
