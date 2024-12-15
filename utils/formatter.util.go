package utils

import (
	"math/big"
)

func BigIntToString(n *big.Int, format int) string {
	return "0x" + n.Text(format)
}

func StringToBigInt(n string, format int) *big.Int {
	formatter := new(big.Int)
	formatter.SetString(n, format)
	return formatter
}
