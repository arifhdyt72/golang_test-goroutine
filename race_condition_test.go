package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Counter : ", counter)
}
