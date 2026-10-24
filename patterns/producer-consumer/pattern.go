package patterns

import (
	"sync"
)

func ProducerConsumer(
	nums []int,
	producers int,
	consumers int,
) []int {
	if producers <= 0 || consumers <= 0 {
		return make([]int, 0)
	}

	ch := make(chan int, producers)

	var muProduser sync.Mutex

	var wgProduser sync.WaitGroup
	wgProduser.Add(producers)

	index := 0
	for range producers {
		go func() {
			defer wgProduser.Done()
			for {
				var num int
				muProduser.Lock()
				if index < len(nums) {
					num = nums[index]
					index++
				} else {
					muProduser.Unlock()
					return
				}
				muProduser.Unlock()

				ch <- num
			}
		}()
	}

	result := make([]int, 0, len(nums))

	var wgConsumer sync.WaitGroup
	var muConsumer sync.Mutex

	wgConsumer.Add(consumers)
	for range consumers {
		go func() {
			defer wgConsumer.Done()

			for num := range ch {
				muConsumer.Lock()
				result = append(result, num*num)
				muConsumer.Unlock()
			}
		}()
	}

	wgProduser.Wait()
	close(ch)
	wgConsumer.Wait()
	return result
}
