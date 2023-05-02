package task

import "strconv"

func FizzBuzz(input []int) []string {
	var result []string

	for _, value := range input {
		switch {
		case value%3 == 0 && value%5 == 0:
			result = append(result, "FizzBuzz")
		case value%3 == 0:
			result = append(result, "Fizz")
		case value%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(value))
		}
	}

	return result
}
