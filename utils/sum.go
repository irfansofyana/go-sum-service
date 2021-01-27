package utils

import (
	"errors"
	"math/big"
)

// Sum returns a + b
func Sum(a string, b string) (*big.Int, error) {
	var bigA, isNumberA = new(big.Int).SetString(a, 10)
	var bigB, isNumberB = new(big.Int).SetString(b, 10)

	if !isNumberA || !isNumberB {
		return nil, errors.New("Cannot perform sum operation: non-number value detected")
	}

	return bigA.Add(bigA, bigB), nil
}
