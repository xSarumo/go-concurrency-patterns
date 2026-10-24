package patterns

import "sync"

func FanIn[T any](inputs ...<-chan T) <-chan T {
	if len(inputs) == 0 {
		result := make(chan T)
		close(result)
		return result
	}

	result := make(chan T, len(inputs))
	var wg sync.WaitGroup

	wg.Add(len(inputs))
	for _, ch := range inputs {
		go func() {
			defer wg.Done()
			for value := range ch {
				result <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func FanOut[T any](input <-chan T, workers int) []<-chan T {
	if workers <= 0 {
		return make([]<-chan T, 0)
	}

	result := make([]<-chan T, 0, workers)

	for range workers {
		ch := make(chan T, 1)
		result = append(result, ch)
		go func() {
			for value := range input {
				ch <- value
			}

			close(ch)
		}()
	}

	return result
}
