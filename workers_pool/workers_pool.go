package workers_pool

import (
	"fmt"
	"sync"
	"time"
)

// Don't change this function
func printNumber(id int) {
	time.Sleep(1 * time.Second)
	fmt.Println(id)
}

func DoPrint(poolSize int) {
	wg := sync.WaitGroup{}

	ch := make(chan int)
	defer close(ch)

	for j := 0; j < poolSize; j++ {
		go func(ch <-chan int) {
			// defer wg.Done()
			for k := range ch {
				printNumber(k)
				wg.Done()
			}
		}(ch)
	}

	for i := 1; i < 11; i++ {
		wg.Add(1)
		ch <- i
	}

	wg.Wait()
}
