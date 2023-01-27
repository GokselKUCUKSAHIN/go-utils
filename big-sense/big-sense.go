package big_sense

import "math/big"

func Add(first *big.Int, second *big.Int) *big.Int {
	return big.NewInt(0).Add(first, second)
}

func AddAll(first *big.Int, rest ...*big.Int) *big.Int {
	sum := big.NewInt(0).Add(big.NewInt(0), first)
	size := len(rest)
	for i := 0; i < size; i++ {
		sum = sum.Add(sum, rest[i])
	}
	return sum
}

func Sub(first *big.Int, second *big.Int) *big.Int {
	return big.NewInt(0).Sub(first, second)
}

func Mul(first *big.Int, second *big.Int) *big.Int {
	return big.NewInt(0).Mul(first, second)
}

func MulAll(first *big.Int, rest ...*big.Int) *big.Int {
	total := big.NewInt(1).Mul(big.NewInt(1), first)
	size := len(rest)
	for i := 0; i < size; i++ {
		total = total.Mul(total, rest[i])
	}
	return total
}

func Div(dividend *big.Int, divisor *big.Int) *big.Int {
	return big.NewInt(0).Div(dividend, divisor)
}
