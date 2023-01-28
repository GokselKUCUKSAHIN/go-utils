package big_sense

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/big"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func ValueOf[T Number](number T) *big.Int {
	int64Value := int64(number)
	return big.NewInt(int64Value)
}

func FromStringBase(number string, base int) (*big.Int, error) {
	if base < 2 {
		return nil, fmt.Errorf("big_sense: base can-not be less than 2. base %v", base)
	}
	n := new(big.Int)
	if n, ok := n.SetString(number, base); ok {
		return n, nil
	}
	return nil, fmt.Errorf("big_sense: '%v' convertion can-not be done", number)
}

func FromString(number string) (*big.Int, error) {
	return FromStringBase(number, 10)
}

func Zero() *big.Int {
	return big.NewInt(0)
}

func One() *big.Int {
	return big.NewInt(1)
}

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
