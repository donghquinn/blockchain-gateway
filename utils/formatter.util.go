package utils

import (
	"fmt"
	"math/big"
)

func BigIntToString(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %x or upper case
}

func StringToBigInt(n string) *big.Int {
	formatter := new(big.Int)
	formatter.SetString(n, 16)
	return formatter
}
