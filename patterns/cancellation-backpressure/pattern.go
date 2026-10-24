package patterns

import "context"

func Forward(
	ctx context.Context,
	input <-chan int,
) <-chan int {
	output := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(output)
				return
			case value, ok := <-input:
				if !ok {
					close(output)
					return
				}
				select {
				case output <- value:
					continue
				case <-ctx.Done():
					close(output)
					return
				}
			}
		}
	}()

	return output
}

func GenerateWithContext(
	ctx context.Context,
	nums ...int,
) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, i := range nums {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				continue
			}
		}
	}()

	return ch
}
