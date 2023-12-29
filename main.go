package main

import (
	"errors"
	"fmt"
	"math"
)

func Add32(left, right int) (int, error) {
	if right > 0 {
		if left > math.MaxInt32-right {
			return 0, errors.New("int32 overflow")
		}
	} else {
		if left < math.MinInt32-right {
			return 0, errors.New("int32 overflow")
		}
	}
	return left + right, nil
}

func main() {

	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	var i int = 2147483647
	fmt.Println(i)
	fmt.Println("i: ", i)
	j, err := Add32(i, 0)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("j: ", j)
	fmt.Println(j > math.MaxInt32)
}
