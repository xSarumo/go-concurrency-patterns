package patterns

import "sync"

func RunLimited(tasks []func(), limit int) {
	if limit == 0 {
		return
	}

	semaphore := make(chan struct{}, limit)

	var wg sync.WaitGroup

	wg.Add(len(tasks))
	for _, task := range tasks {
		semaphore <- struct{}{}
		go func() {
			defer wg.Done()
			task()
			_ = <-semaphore
		}()
	}

	wg.Wait()
}
