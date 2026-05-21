package main

import (
	"errors"
	"fmt"
)

type Number float64

func Add(a, b Number) Number {
	return a + b
}

func Subtract(a, b Number) Number {
	return a - b
}

func Multiply(a, b Number) Number {
	return a * b
}

func Divide(a, b Number) (Number, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func main() {
	x, y := Number(10), Number(3)
	fmt.Printf("%.0f + %.0f = %.0f\n", x, y, Add(x, y))
	fmt.Printf("%.0f - %.0f = %.0f\n", x, y, Subtract(x, y))
	fmt.Printf("%.0f x %.0f = %.0f\n", x, y, Multiply(x, y))

	result, err := Divide(x, y)
	if err != nil {
		fmt.Println("除零错误:", err)
	} else {
		fmt.Printf("%.0f / %.0f = %.2f\n", x, y, result)
	}
}
