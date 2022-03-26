package DateDiff

import (
	"errors"
	"fmt"
)

func UintParse(x string) (uint, error) {
	if len(x) == 0 {
		return 0, errors.New("empty input is invalid")
	}
	y := uint(0)
	n := len(x)
	for i := n - 1; i >= 0; i-- {
		// Convert character to uint. If x[i] is not valid, throw error early
		digit, isOkay := digitParse[x[i]]
		if !isOkay {
			return 0, fmt.Errorf("character %s, at postition %d of input is not a valid Base10 digit", string(x[i]), i)
		}
		y += digit * pow10(uint(n-i-1))
	}
	return y, nil
}

func pow10(x uint) uint {
	//return uint(tenPower[x])
	return uint(tenPower32[x/32] * tenPower[x%32])
}

var tenPower = []float64{1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15,
	1e16, 1e17, 1e18, 1e19, 1e20, 1e21, 1e22, 1e23, 1e24, 1e25, 1e26, 1e27, 1e28, 1e29, 1e30, 1e31}
var tenPower32 = []float64{1e00, 1e32, 1e64, 1e96}
var digitParse = map[uint8]uint{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}
