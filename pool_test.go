package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	group := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}

	pool.Put("Arif")
	pool.Put("Hidayat")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Done Pool Test")
}
