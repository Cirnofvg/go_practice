package chanpatterns

import "sync"

func Merge(chIn ...<-chan int) <-chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(chIn))

	for _, ch := range chIn {
		if ch == nil {
			continue
		}
		defer wg.Done()
		go func() {
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Split(chIn <-chan int, split int) []<-chan int {
	if chIn == nil {
		return nil
	}

	outPutChs := make([]chan int, split)
	for i := 0; i < split; i++ {
		outPutChs[i] = make(chan int)
	}

	go func() {
		idx := 0
		for v := range chIn {
			outPutChs[idx] <- v
			idx = (idx + 1) % split
		}

		for _, v := range outPutChs {
			close(v)
		}
	}()

	resChs := make([]<-chan int, 2)
	for i := 0; i < split; i++ {
		resChs[i] = outPutChs[i]
	}

	return resChs
}

func Tee(chIn <-chan int, split int) []<-chan int {
	putChs := make([]chan int, split)
	for i := 0; i < split; i++ {
		putChs[i] = make(chan int)
	}

	go func() {
		for v := range chIn {
			for _, ch := range putChs {
				ch <- v
			}
		}

		for _, v := range putChs {
			close(v)
		}
	}()

	outChs := make([]<-chan int, split)
	for i := 0; i < split; i++ {
		outChs[i] = putChs[i]
	}

	return outChs
}

func Transform(chIn <-chan int, action func(int) int) <-chan int {
	chOut := make(chan int)
	go func() {
		defer close(chOut)
		for v := range chIn {
			chOut <- action(v)
		}
	}()

	return chOut
}

func Filter(chIn <-chan int, predicate func(int) bool) <-chan int {
	chOut := make(chan int)

	go func() {
		defer close(chOut)
		for v := range chIn {
			if predicate(v) {
				chOut <- v
			}
		}
	}()

	return chOut
}
