package task

import (
	"reflect"
	"testing"
)

func TestEmptyFizzBuzzInput(t *testing.T) {
	empty_slice := make([]int, 0)
	result := FizzBuzz(empty_slice)

	if len(result) != 0 {
		t.Error("Output is not empty")
	}
}

func TestFizzBuzz(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 15}
	expected_result := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "FizzBuzz"}

	fizz := FizzBuzz(input)

	if len(fizz) != len(expected_result) {
		t.Error("Invalid result size")
	}

	// TODO: Use github.com/google/go-cmp
	// go-cmp is intended to be a more powerful and safer alternative to reflect.DeepEqual for comparing whether two values are semantically equal.
	if !reflect.DeepEqual(fizz, expected_result) {
		t.Error("Invalid output result")
	}
}
