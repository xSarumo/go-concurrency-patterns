package patterns

import (
	"errors"
	"sync"
)

type Nums interface {
	~int | ~int16 | ~int32 | ~int64 | ~uint | ~uint32 | ~uint16 | ~uint64
}

func SquareNums[T Nums](nums []T, workers int) ([]T, error) {
	if len(nums) == 0 || workers <= 0 {
		return nil, errors.New("nums is empty or workers less or equal zero!")
	}

	channel := make(chan T, len(nums)/2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range nums {
			channel <- num
		}
		close(channel)
	}()

	result := make([]T, 0, len(nums))

	var mu sync.Mutex
	wg.Add(workers)
	for _ = range workers {
		go func() {
			defer wg.Done()
			for value := range channel {
				value = value * value

				mu.Lock()
				result = append(result, value)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	return result, nil
}
