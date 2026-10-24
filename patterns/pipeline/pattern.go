package patterns

import "fmt"

func Generate(nums ...int) <-chan int {
	output := make(chan int)

	go func() {
		for _, num := range nums {
			output <- num
		}
		close(output)
	}()

	return output
}

func Square(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		for num := range input {
			output <- num * num
		}
		close(output)
	}()

	return output
}

func FilterEven(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		for num := range input {
			if num%2 == 0 {
				output <- num
			}
		}
		close(output)
	}()

	return output
}

func Pipeline() {
	numbers := Generate(1, 2, 3, 4, 5)
	squared := Square(numbers)
	filtered := FilterEven(squared)

	for value := range filtered {
		fmt.Println(value)
	}
}
