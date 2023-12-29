package reverseint

import (
	"errors"
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == 0 || x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}

	isNegative := x < 0

	if isNegative {
		x = -x
	}

	digitsCount := len(fmt.Sprint(x))
	digitsReversed := make([]int, 0, digitsCount)

	for x > 0 {
		digitsReversed = append(digitsReversed, x%10)
		x /= 10
	}

	var err error = nil
	result := 0
	multiplier := 1
	for i := 0; i < digitsCount-1; i++ {
		multiplier, err = Mul32(multiplier, 10)
		if err != nil {
			fmt.Println("err: ", err)
			return 0
		}
	}

	for _, v := range digitsReversed {
		m := 0
		m, err = Mul32(multiplier, v)
		if err != nil {
			fmt.Println("err: ", err)
			return 0
		}
		result, err = Add32(result, m)
		if err != nil {
			fmt.Println("err: ", err)
			return 0
		}
		multiplier /= 10
	}

	if isNegative {
		result = -result
	}

	return result
}

func Add32(left int, right int) (int, error) {
	if right > 0 {
		if left > math.MaxInt32-right {
			return 0, errors.New("int32 overflow during add operation")
		}
	} else {
		if left < math.MinInt32-right {
			return 0, errors.New("int32 overflow during add operation")
		}
	}
	return left + right, nil
}

func Mul32(left int, right int) (int, error) {
	if right > 0 {
		if left > math.MaxInt32*right {
			return 0, errors.New("int32 overflow during multiply operation")
		}
	} else {
		if left < math.MinInt32*right {
			return 0, errors.New("int32 overflow during multiply operation")
		}
	}
	return left * right, nil
}
