package calculator

import (
	"errors"
	"fmt"
)

func Addition(nums ...int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("need minimum two numbers to perform addition")
	}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum, nil
}

func Multiply(nums ...int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("need minimum two numbers to perform multiplication")
	}

	product := 1

	for _, num := range nums {
		product *= num
	}

	return product, nil
}

func Division(num1 int, num2 int) (int, error) {

	if num2 == 0 {
		return 0, fmt.Errorf("can not divide number %d with zero", num1)
	}

	return num1 / num2, nil
}
