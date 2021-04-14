package main

import (
	"sync"
	"testing"
)

// Make a counter which is safe to use concurrently

// Strategy:
// 1. Make a single-threaded counter
// 2. Find bugs
// 3. Make it concurrent

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d wanted %d", got.Value(), want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safe concurrently", func(t *testing.T) {
		counter := Counter{}
		var wg sync.WaitGroup

		totalInc := 100
		wg.Add(totalInc)

		for i := 0; i < totalInc; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, &counter, totalInc)
	})
}
