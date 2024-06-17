package main

import (
	"sync"
	"testing"
)

func BenchmarkSampleFunction(b *testing.B) {
	var wg sync.WaitGroup

	for i := 0; i < b.N; i++ {
		for i := 1; i <= 5; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				TestApi()
			}()
		}

		wg.Wait()
	}

}
