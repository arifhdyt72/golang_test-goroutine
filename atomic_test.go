package golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Total Counter", x)
}
