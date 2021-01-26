package utils

import (
	"errors"
	"math/big"
)

func Sum(a string, b string) (string, error) {
	var bigA, isNumberA = new(big.Int).SetString(a, 10)
	var bigB, isNumberB = new(big.Int).SetString(b, 10)

	if (!isNumberA || !isNumberB) {
		return "", errors.New("Cannot perform sum operation: non-number value detected")
	}

	return bigA.Add(bigA, bigB).Text(10), nil
}