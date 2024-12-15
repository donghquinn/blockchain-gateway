package utils

import (
	"fmt"
	"math/big"
)

func ToHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %x or upper case
}

func BigintToHex(n string) *big.Int {
	formatter := new(big.Int)
	formatter.SetString(n, 16)
	return formatter
}
